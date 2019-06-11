package store

import (
	"github.com/influxdata/influxdb/models"
	"gitlab.bearstech.com/factory/gyumao/rule"
	"gitlab.bearstech.com/factory/gyumao/timeline"
)

type Store struct {
	store map[Key]*timeline.Timeline
}

func New() *Store {
	return &Store{
		store: make(map[Key]*timeline.Timeline)
	}
}

func (s *Store) Append(r *rule.Rule, point models.Point) {
	k := BuildKey(r, point)
	_, ok := s.store[k]
	if !ok {
		s.store[k] = timeline.New(3) // FIXME hardcoded length
	}
	s.store.Set(point.Time(), point)
}
