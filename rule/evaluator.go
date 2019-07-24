package rule

import (
	_expr "github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/influxdata/influxdb/models"
)

type Evaluator interface {
	Eval(models.Point) (bool, error)
}

type ExprEvaluator struct {
	prog *vm.Program
}

func NewExprEvaluator(expr string) (*ExprEvaluator, error) {
	prog, err := _expr.Compile(expr)
	if err != nil {
		return nil, err
	}
	return &ExprEvaluator{
		prog: prog,
	}, nil
}

func (e *ExprEvaluator) Eval(point models.Point) (bool, error) {
	// TODO
	return false, nil
}
