package deadman

import (
	"github.com/willf/bitset"
)

// DeadRegistry stores the dead and the living
type DeadRegistry struct {
	bitset *bitset.BitSet
}

// NewDeadRegistry return DeadRegistry with a fixed size
func NewDeadRegistry(size uint) *DeadRegistry {
	return &DeadRegistry{
		bitset: bitset.New(size),
	}
}

// Alive set ranked buddy as alive
func (d *DeadRegistry) Alive(rank uint) *DeadRegistry {
	d.bitset.Set(rank)
	return d
}

// Reset all the buddies to alive
func (d *DeadRegistry) Reset() *DeadRegistry {
	d.bitset.ClearAll()
	return d
}

// Or is a boolean operation on two DeadRegistries
func (d *DeadRegistry) Or(d2 *DeadRegistry) *DeadRegistry {
	return &DeadRegistry{
		bitset: d.bitset.Union(d2.bitset),
	}
}

// DeadIterator over dead buddies
type DeadIterator struct {
	registry *DeadRegistry
	cpt      uint
}

// Next iterate for the next dead
// for n, ok := i.Next(); ok; n, ok = i.Next() {}
func (di *DeadIterator) Next() (uint, bool) {
	if di.cpt > di.registry.bitset.Len() {
		return 0, false
	}
	rank, ok := di.registry.bitset.NextClear(di.cpt)
	di.cpt = rank + 1
	return rank, ok
}

// DeadIterator returns a DeadIterator
func (d *DeadRegistry) DeadIterator() *DeadIterator {
	return &DeadIterator{
		registry: d,
	}
}
