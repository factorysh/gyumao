package rule

import (
	"testing"
	"time"

	"github.com/influxdata/influxdb/models"
	"github.com/stretchr/testify/assert"
)

func TestRules(t *testing.T) {
	rules := New()
	match := 0
	rules.Append("http", Tags{
		"status":   "200",
		"hostname": "bob-42"}, func(p models.Point) error {
		match++
		return nil
	})
	point, err := models.NewPoint(
		"http",
		models.NewTags(Tags{"status": "200"}),
		models.Fields{"size": 42},
		time.Now())
	assert.NoError(t, err)
	rules.Filter(point)
	assert.Equal(t, 0, match)

	point, err = models.NewPoint(
		"http",
		models.NewTags(Tags{"status": "404"}),
		models.Fields{"size": 42},
		time.Now())
	assert.NoError(t, err)
	rules.Filter(point)
	assert.Equal(t, 0, match)

	/* FIXME
	point, err = models.NewPoint(
		"http",
		models.NewTags(Tags{"status": "200", "hostname": "bob-42"}),
		models.Fields{"size": 42},
		time.Now())
	assert.NoError(t, err)
	rules.Filter(point)

	assert.Equal(t, 1, match)
	*/
}
