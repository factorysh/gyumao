package point

import (
	"github.com/factorysh/gyumao/rule"
	"github.com/influxdata/influxdb/models"
)

// Point is a measurement point and its rule
type Point struct {
	point models.Point
	rule  *rule.Rule
	name  string
}

func New(point models.Point, rule *rule.Rule) *Point {
	return &Point{
		point: point,
		rule:  rule,
	}
}

// Name of the point
func (p *Point) Name() string {
	if p.name == "" {
		p.name = p.Rule().NamePoint(p.point)
	}
	return p.name
}

func (p *Point) Rule() *rule.Rule {
	return p.rule
}

func (p *Point) Point() models.Point {
	return p.point
}
