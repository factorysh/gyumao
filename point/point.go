package point

import (
	"bytes"

	"github.com/factorysh/gyumao/rule"
	"github.com/influxdata/influxdb/models"
)

// Point is a measurement point and its rule
type Point struct {
	point models.Point
	rule  *rule.Rule
}

// Name of the point
func (p *Point) Name() string {
	b := bytes.NewBuffer(p.point.Name())
	for _, k := range p.rule.GroupBy {
		b.WriteRune(',')
		b.WriteString(k)
		b.WriteRune('=')
		b.WriteString(p.point.Tags().GetString(k))
	}
	return b.String()
}
