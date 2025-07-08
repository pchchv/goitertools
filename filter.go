package goitertools

// FilterFn represents the `filterIterator` function.
type FilterFn[T any] func(v T) bool

// filterIterator allows filtering of an `Iterator[T]`.
type filterIterator[T any, I Iterator[T], MAP any] struct {
	iterator I
	fn       FilterFn[T]
}
