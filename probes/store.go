package probes

type Store interface {
	// Exists return true if the key is already registered
	Exists(key string) bool
	// Discover a new key
	Discover(key string)
	// Validate
	Validate(key string)
}
