package store

import (
	"github.com/factorysh/gyumao/timeline"
	"github.com/influxdata/influxdb/models"
)

// Store stores timeline.Timeline
type Store struct {
	store map[Key]*timeline.Timeline
}

// New returns a new Store
func New() *Store {
	return &Store{
		store: make(map[Key]*timeline.Timeline),
	}
}

// Append a new Influxdb point, with the names of the keys
func (s *Store) Append(keys []string, point models.Point) {
	k := BuildKey(keys, point)
	_, ok := s.store[k]
	if !ok {
		s.store[k] = timeline.New(3) // FIXME hardcoded length
	}
	s.store[k].Set(point.Time(), point)
}
