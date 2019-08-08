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
	v := func(p models.Point) error {
		match++
		return nil
	}
	err = rules.Visit(point, nil, v)
	assert.Equal(t, 1, match)

	point, err = models.NewPoint(
		"http",
		models.NewTags(map[string]string{"status": "404"}),
		models.Fields{"size": 2},
		time.Now())
	assert.NoError(t, err)

	match = 0
	err = rules.Visit(point, nil, v)
	assert.Equal(t, 0, match)
}

func TestExpr(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	rules, err := FromConfig(&config.Config{
		Rules: []*config.Rule{
			&config.Rule{
				Measurement: "http",
				TagsPass: map[string][]string{
					"status": []string{"200"},
				},
				Expr: "size > 50",
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
	v := func(p models.Point) error {
		match++
		return nil
	}
	err = rules.Visit(point, nil, v)
	assert.Equal(t, 0, match)

	point, err = models.NewPoint(
		"http",
		models.NewTags(map[string]string{"status": "200"}),
		models.Fields{"size": 52},
		time.Now())
	assert.NoError(t, err)

	match = 0
	err = rules.Visit(point, nil, v)
	assert.Equal(t, 1, match)
}
