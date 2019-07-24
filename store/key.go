package store

import (
	"bytes"

	"github.com/influxdata/influxdb/models"
)

// Key is a combined key, for storing timeline.Timeline in Store
type Key string

// BuildKey returns Key from a list of tag name, and an Influxdb point
func BuildKey(keys []string, point models.Point) Key {
	b := bytes.Buffer{}
	b.Write(point.Name())
	for _, key := range keys {
		b.WriteRune(',')
		b.WriteString(key)
		b.WriteRune('=')
		b.Write(point.Tags().Get([]byte(key)))
	}
	return Key(b.String())
}
