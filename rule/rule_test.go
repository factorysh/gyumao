package rule

import (
	"testing"
	"time"

	"github.com/factorysh/gyumao/config"
	"github.com/influxdata/influxdb/models"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRules(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	rules, err := FromRules(
		[]*config.Rule{
			&config.Rule{
				Measurement: "http",
				TagsPass: map[string][]string{
					"status": []string{"200"},
				},
			},
		}...)
	assert.NoError(t, err)

	point, err := models.NewPoint(
		"http",
		models.NewTags(map[string]string{"status": "200"}),
		models.Fields{"size": 42},
		time.Now())
	assert.NoError(t, err)

	match := 0
	v := func(r *Rule, p models.Point) error {
		match++
		return nil
	}
	probes := &YesProbes{}
	err = rules.Filter(point, probes, v)
	assert.Equal(t, 1, match)

	point, err = models.NewPoint(
		"http",
		models.NewTags(map[string]string{"status": "404"}),
		models.Fields{"size": 2},
		time.Now())
	assert.NoError(t, err)

	match = 0
	err = rules.Filter(point, probes, v)
	assert.Equal(t, 0, match)
}

type YesProbes struct{}

func (y *YesProbes) Exist(key string) bool { return true }

func (y *YesProbes) Unknown() []string { return make([]string, 0) }

func (y *YesProbes) Keys() []string { return make([]string, 0) }
