package consumer

import (
	"fmt"

	"github.com/factorysh/gyumao/point"
	"github.com/factorysh/gyumao/timeline"
)

type Consumer struct {
	global    map[string]interface{}
	timelines map[string]*timeline.Timeline
}

func NewConsumer() *Consumer {
	return &Consumer{
		global: make(map[string]interface{}),
	}
}

func (c *Consumer) Consume(point *point.Point) error {
	ok, err := point.Rule().Evaluator.Eval(point.Point(), c.global)
	if err != nil {
		return err
	}
	fmt.Println(ok)
	return nil
}
