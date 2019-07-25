package rule

import (
	"testing"
	"time"

	"github.com/factorysh/gyumao/config"
	"github.com/influxdata/influxdb/models"
	"github.com/stretchr/testify/assert"
)

func TestRules(t *testing.T) {
	rules, err := FromConfig(&config.Config{
		Rules: []*config.Rule{
			&config.Rule{
				Measurement: "http",
				TagsPass: map[string][]string{
					"status": []string{"200"},
				},
			},
		},
	})
	assert.NoError(t, err)

	point, err := models.NewPoint(
		"http",
		models.NewTags(map[string]string{"status": "200"}),
		models.Fields{"size": 42},
		time.Now())
	assert.NoError(t, err)

	match := 0
	err = rules.Visit(point, func(p models.Point) error {
		match++
		return nil
	})
	assert.Equal(t, 1, match)

}
