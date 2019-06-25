package store

import (
	"github.com/influxdata/influxdb/models"
	"gitlab.bearstech.com/factory/gyumao/timeline"
)

type Store struct {
	store map[Key]*timeline.Timeline
}

func New() *Store {
	return &Store{
		store: make(map[Key]*timeline.Timeline),
	}
}

func (s *Store) Append(keys []string, point models.Point) {
	k := BuildKey(keys, point)
	_, ok := s.store[k]
	if !ok {
		s.store[k] = timeline.New(3) // FIXME hardcoded length
	}
	s.store[k].Set(point.Time(), point)
}
