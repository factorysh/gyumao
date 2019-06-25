package timeline

import (
	"sync"
	"time"
)

type Timeline struct {
	size  int
	store map[time.Time]interface{}
	keys  []time.Time
	lock  sync.Mutex
}

func New(size int) *Timeline {
	return &Timeline{
		size:  size,
		store: make(map[time.Time]interface{}),
		keys:  make([]time.Time, 0),
		lock:  sync.Mutex{},
	}
}

func (t *Timeline) Copy() *Timeline {
	t.lock.Lock()
	defer t.lock.Unlock()
	tt := New(t.size)
	tt.keys = make([]time.Time, t.Lenght())
	for i, key := range t.keys {
		tt.keys[i] = key
		tt.store[key] = t.store[key]
	}
	return tt
}

func (t *Timeline) Set(ts time.Time, v interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if len(t.keys) >= t.size { //full
		first := t.keys[0]
		delete(t.store, first)
		t.keys = t.keys[1:]
	}
	if len(t.keys) == 0 {
		t.keys = []time.Time{ts}
		t.store[ts] = v
		return
	}
	if t.keys[len(t.keys)-1].After(ts) {
		// sort is fucked
		panic("please sort, or patch the code")
	} else {
		t.keys = append(t.keys, ts)
		t.store[ts] = v
	}
}

func (t *Timeline) First() interface{} {
	return t.store[t.keys[0]]
}

func (t *Timeline) Last() interface{} {
	return t.store[t.keys[len(t.keys)-1]]
}

func (t *Timeline) Lenght() int {
	return len(t.keys)
}

func (t *Timeline) Duration(d time.Duration) *Timeline {
	tt := t.Copy()
	since := time.Now().Add(d)
	var i int
	for _, k := range t.keys {
		i++
		if k.Before(since) {
			delete(tt.store, k)
		} else {
			break
		}
	}
	tt.keys = tt.keys[i:]
	return tt
}
