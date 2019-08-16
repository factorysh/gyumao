package rule

import (
	"bytes"

	"github.com/factorysh/gyumao/config"
	"github.com/factorysh/gyumao/evaluator"
	 "github.com/factorysh/gyumao/evaluator/expr"
	_probes "github.com/factorysh/gyumao/probes"
	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/telegraf/filter"
	log "github.com/sirupsen/logrus"
)

type Rule struct {
	Measurement string
	TagsPass    Tags
	TagsExclude Tags
	GroupBy     []string // FIXME ensure keys are sorted
	Evaluator   evaluator.Evaluator
}

type Rules map[string][]*Rule
type Tags map[string]filter.Filter

func tags(in map[string][]string) (Tags, error) {
	t := make(map[string]filter.Filter)
	for k, v := range in {
		var err error
		t[k], err = filter.Compile(v)
		if err != nil {
			return nil, err
		}
	}
	return Tags(t), nil
}

// New Rules
func New() Rules {
	return make(map[string][]*Rule)
}

// FromRules build a Rules from a list of Rule
func FromRules(lRules ...*config.Rule) (Rules, error) {
	rules := New()
	for _, rule := range lRules {
		l := log.WithField("rule", rule)
		var e evaluator.Evaluator
		var err error
		if rule.Expr != "" {
			e, err = expr.New(rule.Expr)
			if err != nil {
				l.WithError(err).Error()
				return nil, err
			}
		} else {
			e = &evaluator.YesEvaluator{}
		}

		tp, err := tags(rule.TagsPass)
		if err != nil {
			l.WithError(err).Error()
			return nil, err
		}
		te, err := tags(rule.TagsExclude)
		if err != nil {
			l.WithError(err).Error()
			return nil, err
		}
		rules.Append(rule.Measurement, &Rule{
			Measurement: rule.Measurement,
			TagsPass:    tp,
			TagsExclude: te,
			Evaluator:   e,
		})
	}
	return rules, nil
}

// Append a new Rule, with its name
func (r Rules) Append(name string, rule *Rule) {
	_, ok := r[name]
	if !ok {
		r[name] = []*Rule{rule}
	} else {
		r[name] = append(r[name], rule)
	}
}

// Filter a point
func (r *Rule) Filter(point models.Point) bool {
	l := log.WithField("point", point)
	for t, filter := range r.TagsPass {
		l = l.WithField("tag name", t)
		tag := []byte(t)
		if !point.HasTag(tag) {
			l.Info("No tag")
			return false
		}
		v := point.Tags().Get(tag)
		if !filter.Match(string(v)) {
			l.WithField("value", string(v)).Info("Don't match")
			return false
		}
	}
	for t, filter := range r.TagsExclude {
		tag := []byte(t)
		if !point.HasTag(tag) {
			return false
		}
		v := point.Tags().Get(tag)
		if filter.Match(string(v)) {
			return false
		}
	}
	return true
}

func (r *Rule) NamePoint(point models.Point) string {
	b := bytes.NewBuffer(point.Name())
	for _, k := range r.GroupBy {
		b.WriteRune(',')
		b.WriteString(k)
		b.WriteRune('=')
		b.WriteString(point.Tags().GetString(k))
	}
	return b.String()
}

// Filter all Rules with a Point and a callback
func (r Rules) Filter(point models.Point, probes _probes.Probes,
	do func(r *Rule, point models.Point) error) error {
	name := string(point.Name())
	rules, ok := r[name]
	if !ok {
		return nil
	}
	for _, rule := range rules {
		if rule.Filter(point) && probes.Exist(rule.NamePoint(point)) {
			err := do(rule, point)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
