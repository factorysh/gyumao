package deadman

import (
	"github.com/willf/bitset"
)

type DeadRegistry struct {
	bitset *bitset.BitSet
}

func New(size uint) *DeadRegistry {
	return &DeadRegistry{
		bitset: bitset.New(size),
	}
}

func (d *DeadRegistry) Alive(rank uint) {
	d.bitset.Set(rank)
}

type DeadIterator struct {
	registry *DeadRegistry
	cpt      uint
}

func (di *DeadIterator) Next() (uint, bool) {
	rank, ok := di.registry.bitset.NextClear(di.cpt)
	if rank == di.cpt {
		return 0, false
	}
	di.cpt = rank
	return di.cpt, ok
}

func (d *DeadRegistry) DeadIterator() *DeadIterator {
	return &DeadIterator{
		registry: d,
	}
}
