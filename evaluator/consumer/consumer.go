package consumer

import (
	"fmt"

	"github.com/factorysh/gyumao/point"
	"github.com/factorysh/gyumao/timeline"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	global    map[string]interface{}
	timelines map[string]*timeline.Timeline
}

func NewConsumer(global map[string]interface{}) *Consumer {
	if global == nil {
		global = make(map[string]interface{})
	}
	return &Consumer{global: global}
}

func (c *Consumer) Consume(point *point.Point) error {
	l := log.WithField("point", point)
	ok, err := point.Rule().Evaluator.Eval(point.Point(), c.global)
	if err != nil {
		l.WithError(err).Error("evaluator/consumer.Consumer#Consume")
		return err
	}
	l.Info("evaluator/consumer.Consumer#Consume")
	fmt.Println(ok)
	return nil
}
