package evaluator

import (
	"github.com/influxdata/influxdb/models"
)

type Evaluator interface {
	Eval(point models.Point, context map[string]interface{}) (bool, error)
}

type YesEvaluator struct {
}

func (y *YesEvaluator) Eval(models.Point, map[string]interface{}) (bool, error) {
	return true, nil
}
