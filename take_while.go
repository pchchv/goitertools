package goitertools

// TakeWhileFn represents the `takeWhileIterator[T]` function.
type TakeWhileFn[T any] func(v T) bool
