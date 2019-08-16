package probes

import "sync"

// Probes is a collection of Probe, with is their contexts
type Probes interface {
	Exist(key string) bool
	Unknown() []string
}

// MapProbes implement basic Probes, with some map
type MapProbes struct {
	probes  map[string]interface{}
	unknown map[string]interface{}
	lock    sync.RWMutex
}

func NewMapProbes() *MapProbes {
	return &MapProbes{
		probes:  make(map[string]interface{}),
		unknown: make(map[string]interface{}),
	}
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
