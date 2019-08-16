package store

import (
	"github.com/armon/go-radix"
	"github.com/factorysh/gyumao/timeline"
	"github.com/influxdata/influxdb/models"
)

// Store stores timeline.Timeline
type Store struct {
	store *radix.Tree
}

// New returns a new Store
func New() *Store {
	return &Store{
		store: radix.New(),
	}
}

// Append a new Influxdb point, with the names of the keys
func (s *Store) Append(keys []string, point models.Point) {
	k := BuildKey(keys, point)
	tl, ok := s.store.Get(k)
	if !ok {
		s.store.Insert(k, timeline.New(3)) // FIXME hardcoded length
	}
	tl.(*timeline.Timeline).Set(point.Time(), point)
}
