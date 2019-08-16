package point

import (
	"fmt"
	"testing"
	"time"

	"github.com/factorysh/gyumao/rule"
	"github.com/influxdata/influxdb/models"
	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	p, err := models.NewPoint("http_status", models.Tags{
		models.NewTag([]byte("dc"), []byte("ilyad-dc4")),
		models.NewTag([]byte("project"), []byte("canary")),
		models.NewTag([]byte("criticity"), []byte("admin")),
	}, models.Fields{
		"time": 42,
	}, time.Now())
	assert.NoError(t, err)
	fmt.Println(p)
	point := &Point{
		point: p,
		rule: &rule.Rule{
			GroupBy: []string{"dc", "project"},
		},
	}
	assert.Equal(t, "http_status,dc=ilyad-dc4,project=canary", point.Name())
}
