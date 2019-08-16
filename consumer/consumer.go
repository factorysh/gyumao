package consumer

import "github.com/factorysh/gyumao/point"

// Consumer consumes points
type Consumer interface {
	Consume(point *point.Point) error
}

// MultiConsumer all the things
type MultiConsumer struct {
	consumers []Consumer
}

// NewMultiConsumer spawn a MultiConsumer
func NewMultiConsumer(consumers ...Consumer) *MultiConsumer {
	return &MultiConsumer{
		consumers: consumers,
	}
}

// Consume consume a point
func (m *MultiConsumer) Consume(point *point.Point) error {
	for _, consumer := range m.consumers {
		err := consumer.Consume(point)
		if err != nil {
			return err
		}
	}
	return nil
}
