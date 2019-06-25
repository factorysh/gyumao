package store

import (
	"testing"
	"time"

	"github.com/influxdata/influxdb/models"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	s := New()
	assert.Len(t, s.store, 0)
	now := time.Now()
	pt, err := models.NewPoint("bob",
		models.NewTags(map[string]string{
			"host":    "bob",
			"service": "httpd",
			"status":  "200",
		}), models.Fields{
			"hits": 342,
		}, now)
	assert.NoError(t, err)
	s.Append([]string{"host", "service"}, pt)
	assert.Len(t, s.store, 1)
}
