package rule

import (
	"github.com/factorysh/gyumao/config"
	"github.com/influxdata/influxdb/models"
)

type Rule struct {
	Measurement string
	Tags        models.Tags
	Keys        []string
	Evaluator   Evaluator
}

type Rules map[string][]*Rule
type Tags map[string]string

func New() Rules {
	return make(map[string][]*Rule)
}

func FromConfig(cnf *config.Config) (Rules, error) {
	rules := New()
	for _, rule := range cnf.Rules {
		_, ok := rules[rule.Measurement]
		if !ok {
			rules[rule.Measurement] = make([]*Rule, 0)
		}
		e, err := NewExprEvaluator(rule.Expr)
		if err != nil {
			return nil, err
		}
		rules[rule.Measurement] = append(rules[rule.Measurement], &Rule{
			Measurement: rule.Measurement,
			Tags:        models.NewTags(rule.TagPass),
			Evaluator:   e,
		})
	}
	return rules, nil
}

func (r Rules) Append(name string, tags Tags, do func(point models.Point) error) {
	_, ok := r[name]
	if !ok {
		r[name] = make([]*Rule, 0)
	}
	r[name] = append(r[name], &Rule{
		Measurement: name,
		Tags:        models.NewTags(tags),
	})
}

func (r Rules) Filter(point models.Point) error {
	rr, ok := r[string(point.Name())]
	if !ok {
		return nil
	}
	for _, rule := range rr {
		if models.CompareTags(point.Tags(), rule.Tags) <= 0 {
		}
	}
	return nil
}
