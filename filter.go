package goitertools

import "github.com/pchchv/express/optionext"

// FilterFn represents the `filterIterator` function.
type FilterFn[T any] func(v T) bool

// filterIterator allows filtering of an `Iterator[T]`.
type filterIterator[T any, I Iterator[T], MAP any] struct {
	iterator I
	fn       FilterFn[T]
}

// Next yields the next value from the iterator that passed the filter function.
func (i *filterIterator[T, I, MAP]) Next() optionext.Option[T] {
	for {
		v := i.iterator.Next()
		if v.IsNone() || !i.fn(v.Unwrap()) {
			return v
		}
	}
}

// Iter is a convenience function that converts the
// `filterIterator` iterator into an `Iterate[T]`.
func (i *filterIterator[T, I, MAP]) Iter() Iterate[T, Iterator[T], MAP] {
	return IterMap[T, Iterator[T], MAP](i)
}
