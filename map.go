package goitertools

import (
	"github.com/pchchv/express/optionext"
	mapext "github.com/pchchv/goitertools/map"
)

// Entry represents a single Map entry.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// mapWrapper is used to transform elements from one type to another.
type mapWrapper[K comparable, V, MAP any] struct {
	m map[K]V
}

// Len returns the underlying map's length.
func (i mapWrapper[K, V, MAP]) Len() int {
	return len(i.m)
}

// Next returns the next transformed element or None if at the end of the iterator.
//
// Warning: This consumes(removes) the map entries as it iterates.
func (i mapWrapper[K, V, MAP]) Next() optionext.Option[Entry[K, V]] {
	for k, v := range i.m {
		delete(i.m, k)
		return optionext.Some(Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return optionext.None[Entry[K, V]]()
}

// Iter is a convenience function that converts the map iterator into an `*Iterate[T]`.
func (i mapWrapper[K, V, MAP]) Iter() Iterate[Entry[K, V], mapWrapper[K, V, MAP], MAP] {
	return IterMap[Entry[K, V], mapWrapper[K, V, MAP], MAP](i)
}

// Retain retains only the elements specified by the function and removes others.
func (i mapWrapper[K, V, MAP]) Retain(fn func(key K, value V) bool) mapWrapper[K, V, MAP] {
	mapext.Retain(i.m, fn)
	return i
}

// WrapMapWithMap creates a new `mapWrapper` for
// use which also specifies a potential future `Map` operation.
func WrapMapWithMap[K comparable, V, MAP any](m map[K]V) mapWrapper[K, V, MAP] {
	return mapWrapper[K, V, MAP]{
		m: m,
	}
}
