package probes

// Probes is a collection of Probe, with is their contexts
type Probes interface {
	Exist(key string) bool
}
