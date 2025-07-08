package goitertools

// FilterFn represents the `filterIterator` function.
type FilterFn[T any] func(v T) bool
