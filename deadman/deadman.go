package deadman

import (
	"sync"

	"github.com/armon/go-radix"
	"github.com/willf/bitset"
)

// DeadRegistry stores the dead and the living
type DeadRegistry struct {
	bitset   *bitset.BitSet
	keys     []string
	keysRank *radix.Tree
	lock     sync.RWMutex
}

// NewDeadRegistry return DeadRegistry from a collection of keys
func NewDeadRegistry(keys []string) *DeadRegistry {
	r := radix.New()
	for i, k := range keys {
		r.Insert(k, uint(i))
	}
	return &DeadRegistry{
		bitset:   bitset.New(uint(len(keys))),
		keys:     keys,
		keysRank: r,
	}
}

// Ghost returns a ghost registry, sharing users
func (d *DeadRegistry) Ghost() *DeadRegistry {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return &DeadRegistry{
		bitset:   bitset.New(uint(len(d.keys))),
		keys:     d.keys,
		keysRank: d.keysRank,
	}
}

// Alive set key as alive
func (d *DeadRegistry) Alive(key string) *DeadRegistry {
	d.lock.RLock()
	defer d.lock.RUnlock()
	rank, ok := d.keysRank.Get(key)
	if ok {
		d.bitset.Set(rank.(uint))
	}
	return d
}

// Reset all the buddies to alive
func (d *DeadRegistry) Reset() *DeadRegistry {
	d.lock.Lock()
	defer d.lock.Unlock()
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
	bitset   *bitset.BitSet
	cpt      uint
}

// Next iterate for the next dead
// for n, ok := i.Next(); ok; n, ok = i.Next() {}
func (di *DeadIterator) Next() (string, bool) {
	if di.cpt > di.bitset.Len() {
		return "", false
	}
	rank, ok := di.bitset.NextClear(di.cpt)
	di.cpt = rank + 1
	return di.registry.keys[rank], ok
}

// DeadIterator returns a DeadIterator
//
// DeadIterator uses a mutex, if you don't iterate all the key,
// the DeadIterator will be stuck
func (d *DeadRegistry) DeadIterator() *DeadIterator {
	return &DeadIterator{
		registry: d,
		bitset:   d.bitset.Clone(),
	}
}
