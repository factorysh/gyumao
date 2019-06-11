package store

import (
	"testing"

	"github.com/influxdata/influxdb/models"
)

func TestStore(t *testing.T) {
	s := New()
	models.NewPoint("bob", models.NewTags())
}
