package rule

import (
	"github.com/influxdata/influxdb/models"
	"gitlab.bearstech.com/factory/gyumao/config"
)

type Rule struct {
	Name string
	Tags models.Tags
	Keys []string
	Do   func(point models.Point) error
}

type Rules map[string][]*Rule
type Tags map[string]string

func New() Rules {
	return make(map[string][]*Rule)
}

func FromConfig(cnf *config.Rules) (Rules, error) {
	rules := New()
	for _, rule := range cnf.Rules {
		_, ok := rules[rule.Name]
		if !ok {
			rules[rule.Name] = make([]*Rule, 0)
		}
		rules[rule.Name] = append(rules[rule.Name], &Rule{
			Name: rule.Name,
			Tags: models.NewTags(rule.Tags),
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
		Name: name,
		Tags: models.NewTags(tags),
		Do:   do,
	})
}

func (r Rules) Filter(point models.Point) error {
	rr, ok := r[string(point.Name())]
	if !ok {
		return nil
	}
	for _, rule := range rr {
		if models.CompareTags(point.Tags(), rule.Tags) <= 0 {
			err := rule.Do(point)
			if err != nil {
				// Crash early
				return err
			}
		}
	}
	return nil
}
