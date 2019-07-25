package rule

import (
	"github.com/factorysh/gyumao/config"
	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/telegraf/filter"
	log "github.com/sirupsen/logrus"
)

type Rule struct {
	Measurement string
	TagsPass    Tags
	TagsExclude Tags
	Keys        []string
	Evaluator   Evaluator
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

// FromConfig build a Rules from a *config.Config
func FromConfig(cnf *config.Config) (Rules, error) {
	rules := New()
	l := log.WithField("config", cnf)
	for _, rule := range cnf.Rules {
		l = l.WithField("rule", *rule)
		var e *ExprEvaluator
		var err error
		if rule.Expr != "" {
			e, err = NewExprEvaluator(rule.Expr)
			if err != nil {
				l.WithError(err).Error()
				return nil, err
			}
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

func (r Rules) Visit(point models.Point, do func(point models.Point) error) error {
	name := string(point.Name())
	rules, ok := r[name]
	if !ok {
		return nil
	}
	for _, rule := range rules {
		for t, filter := range rule.TagsPass {
			tag := []byte(t)
			if !point.HasTag(tag) {
				continue
			}
			v := point.Tags().Get(tag)
			if !filter.Match(string(v)) {
				continue
			}
		}
		for t, filter := range rule.TagsExclude {
			tag := []byte(t)
			if !point.HasTag(tag) {
				continue
			}
			v := point.Tags().Get(tag)
			if filter.Match(string(v)) {
				continue
			}
		}
		err := do(point)
		if err != nil {
			return err
		}
	}
	return nil
}
