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
	rules.Append("http", Tags{"status": "200"}, func(p models.Point) error {
		match += 1
		return nil
	})
	point, err := models.NewPoint("http",
		models.NewTags(map[string]string{"status": "200"}),
		map[string]interface{}{"size": 42},
		time.Now())
	assert.NoError(t, err)
	rules.Filter(point)

	assert.Equal(t, 1, match)
}
