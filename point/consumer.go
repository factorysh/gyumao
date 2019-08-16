package point

// Consumer consumes point.Point
type Consumer interface {
	Consume(point *Point) error
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
func (m *MultiConsumer) Consume(point *Point) error {
	for _, consumer := range m.consumers {
		err := consumer.Consume(point)
		if err != nil {
			return err
		}
	}
	return nil
}
