package goitertools

// Entry represents a single Map entry.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
