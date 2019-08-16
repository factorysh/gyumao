package deadman

import "github.com/factorysh/gyumao/point"

type Consumer struct {
	genealogy *Genealogy
}

func (c *Consumer) Consumer(point *point.Point) error {
	c.genealogy.Current().Alive(point.Name())
	return nil
}
