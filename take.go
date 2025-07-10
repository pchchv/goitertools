package goitertools

// takeIterator is an iterator that only iterates over n elements.
type takeIterator[T any, I Iterator[T], MAP any] struct {
	iterator I
	limit    int
}
