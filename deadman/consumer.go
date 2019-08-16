package deadman

import "github.com/factorysh/gyumao/point"

type Consumer struct {
	genealogy *Genealogy
}

func NewConsumer(genealogy *Genealogy) *Consumer {
	return &Consumer{
		genealogy: genealogy,
	}
}

func (c *Consumer) Consume(point *point.Point) error {
	c.genealogy.Current().Alive(point.Name())
	return nil
}
