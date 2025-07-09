package goitertools

// chainIterator takes two iterators and creates a new iterator over both in sequence.
type chainIterator[T any, FI Iterator[T], SI Iterator[T], MAP any] struct {
	current FI
	next    SI
	flipped bool
}
