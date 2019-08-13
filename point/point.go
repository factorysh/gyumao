package point

import (
	"github.com/factorysh/gyumao/rule"
	"github.com/influxdata/influxdb/models"
)

// Point is a measurement point and its rule
type Point struct {
	point models.Point
	rule  *rule.Rule
}

// Consumer consumes points
type Consumer interface {
	Consume(point *Point) error
}
