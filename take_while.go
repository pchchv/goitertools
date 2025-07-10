package goitertools

// TakeWhileFn represents the `takeWhileIterator[T]` function.
type TakeWhileFn[T any] func(v T) bool

// takeWhileIterator is an iterator that iterates over elements until
// the function return false or end of the iterator
// (whichever happens first).
type takeWhileIterator[T any, I Iterator[T], MAP any] struct {
	iterator Iterator[T]
	fn       TakeWhileFn[T]
}
