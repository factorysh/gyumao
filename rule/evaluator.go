package rule

import (
	"fmt"

	"github.com/antonmedv/expr"
	_expr "github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/influxdata/influxdb/models"
	log "github.com/sirupsen/logrus"
)

type Evaluator interface {
	Eval(point models.Point, context map[string]interface{}) (bool, error)
}

type YesEvaluator struct {
}

func (y *YesEvaluator) Eval(models.Point, map[string]interface{}) (bool, error) {
	return true, nil
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

func (e *ExprEvaluator) Eval(point models.Point, context map[string]interface{}) (bool, error) {
	l := log.WithFields(log.Fields{
		"Point":      point,
		"Expression": e.prog.Source.Content(),
		"Context":    context,
	})
	fields, err := point.Fields()
	if err != nil {
		l.WithError(err).Error()
		return false, err
	}

	l = l.WithField("env", fields)
	l.Info()
	if context == nil {
		context = make(map[string]interface{})
	}
	for k, v := range fields {
		context[k] = v
	}

	out, err := expr.Run(e.prog, context)
	if err != nil {
		fmt.Println(err)
		l.WithError(err).Error()
		return false, err
	}
	resp, ok := out.(bool)
	if !ok {
		err = fmt.Errorf("Expr returns bad type: %p", out)
		log.WithError(err).Error()
		return false, err
	}
	l.Info()
	return resp, nil
}
