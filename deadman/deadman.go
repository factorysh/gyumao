package deadman

import (
	"fmt"
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
func (d *DeadRegistry) DeadIterator() *DeadIterator {
	return &DeadIterator{
		registry: d,
		bitset:   d.bitset.Clone(),
	}
}

// Add more keys
func (d *DeadRegistry) Add(keys ...string) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	start := len(d.keys)
	d.keys = append(d.keys, keys...)
	for i, k := range keys {
		_, ok := d.keysRank.Get(k)
		if ok {
			return fmt.Errorf("Duplicate key : %s", k)
		}
		d.keysRank.Insert(k, uint(start+i))
	}
	return nil
}

// Remove keys
func (d *DeadRegistry) Remove(keys ...string) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	ids := make([]uint, len(keys))
	keysRank := make(map[string]uint)
	for i, k := range keys {
		v, ok := d.keysRank.Get(k)
		if !ok {
			return fmt.Errorf("Unknown key : %s", k)
		}
		ids[i] = v.(uint)
		keysRank[k] = v.(uint)
	}
	newLength := len(d.keys) - len(keys)
	newBitset := bitset.New(uint(newLength))
	newKeys := make([]string, newLength)
	newKeysRank := radix.New()
	rank := 0
	for i, k := range d.keys {
		_, ok := keysRank[k]
		if !ok {
			newKeys[rank] = k
			newKeysRank.Insert(k, uint(rank))
			newBitset.SetTo(uint(rank), d.bitset.Test(uint(i)))
			rank++
		}
	}
	d.bitset = newBitset
	d.keys = newKeys
	d.keysRank = newKeysRank
	return nil
}
