package deadman

import (
	"github.com/willf/bitset"
)

type DeadRegistry struct {
	bitset *bitset.BitSet
}

func NewDeadRegistry(size uint) *DeadRegistry {
	return &DeadRegistry{
		bitset: bitset.New(size),
	}
}

func (d *DeadRegistry) Alive(rank uint) *DeadRegistry {
	d.bitset.Set(rank)
	return d
}

func (d *DeadRegistry) Reset() *DeadRegistry {
	d.bitset.ClearAll()
	return d
}

func (d1 *DeadRegistry) Or(d2 *DeadRegistry) *DeadRegistry {
	return &DeadRegistry{
		bitset: d1.bitset.Union(d2.bitset),
	}
}

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

func (d *DeadRegistry) DeadIterator() *DeadIterator {
	return &DeadIterator{
		registry: d,
	}
}
