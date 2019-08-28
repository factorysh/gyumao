package deadman

import (
	"github.com/factorysh/gyumao/point"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	genealogy *Genealogy
}

func NewConsumer(genealogy *Genealogy) *Consumer {
	return &Consumer{
		genealogy: genealogy,
	}
}

func (c *Consumer) Consume(point *point.Point) error {
	log.WithField("point", point).Info("deadman.Consumer#Consume")
	c.genealogy.Current().Alive(point.Name())
	return nil
}
