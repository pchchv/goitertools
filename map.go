package goitertools

// Entry represents a single Map entry.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// mapWrapper is used to transform elements from one type to another.
type mapWrapper[K comparable, V, MAP any] struct {
	m map[K]V
}
