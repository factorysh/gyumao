package store

import (
	"bytes"

	"github.com/influxdata/influxdb/models"
)

type Key string

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
