package probes

import "sync"

var ProbesPlugin map[string]ProbesFactory

// Probes is a collection of Probe, with is their contexts
type Probes interface {
	Exist(key string) bool
	Unknown() []string
	Keys() []string
}

type ProbesFactory func(map[string]interface{}) (Probes, error)

// MapProbes implement basic Probes, with some map
type MapProbes struct {
	probes  map[string]interface{}
	unknown map[string]interface{}
	keys    []string
	lock    sync.RWMutex
}

func NewMapProbes(keys []string) *MapProbes {
	m := &MapProbes{
		keys:    keys,
		probes:  make(map[string]interface{}),
		unknown: make(map[string]interface{}),
	}
	for _, k := range keys {
		m.probes[k] = new(interface{})
	}
	return m
}

// Exist is this key exists ?
func (m *MapProbes) Exist(key string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.probes[key]
	if !ok {
		m.unknown[key] = new(interface{})
	}
	return ok
}

// Unknown things
func (m *MapProbes) Unknown() []string {
	m.lock.Lock()
	defer m.lock.Unlock()
	u := make([]string, len(m.unknown))
	cpt := 0
	for k := range m.unknown {
		u[cpt] = k
		cpt++
	}
	return u
}

func (m *MapProbes) Keys() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.keys
}
