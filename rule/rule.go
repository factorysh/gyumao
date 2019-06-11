package rule

import (
	"github.com/influxdata/influxdb/models"
)

type Rule struct {
	Name string
	Tags models.Tags
	Do   func(point models.Point) error
}

type Rules map[string][]*Rule
type Tags map[string]string

func New() Rules {
	return make(map[string][]*Rule)
}

func (r Rules) Append(name string, tags map[string]string, do func(point models.Point) error) {
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
		if models.CompareTags(point.Tags(), rule.Tags) >= 0 {
			err := rule.Do(point)
			if err != nil {
				// Crash early
				return err
			}
		}
	}
	return nil
}
