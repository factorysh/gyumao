package store

import (
	"bytes"

	"github.com/influxdata/influxdb/models"
	"gitlab.bearstech.com/factory/gyumao/rule"
)

type Key string

func BuildKey(r *rule.Rule, point models.Point) Key {
	b := bytes.Buffer{}
	b.Write(point.Name())
	for _, key := range r.Keys {
		b.WriteRune(',')
		b.WriteString(key)
		b.WriteRune('=')
		b.Write(point.Tags().Get([]byte(key)))
	}
	return b.String()
}
