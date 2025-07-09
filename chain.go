package goitertools

import "github.com/pchchv/express/optionext"

// chainIterator takes two iterators and creates a new iterator over both in sequence.
type chainIterator[T any, FI Iterator[T], SI Iterator[T], MAP any] struct {
	current FI
	next    SI
	flipped bool
}

// Next returns the next value from the first iterator until exhausted and then the second.
func (i *chainIterator[T, FI, SI, MAP]) Next() optionext.Option[T] {
	for {
		if i.flipped {
			return i.next.Next()
		}

		if v := i.current.Next(); v.IsSome() {
			return v
		}

		i.flipped = true
	}
}

// Iter is a convenience function that converts the chainIterator iterator into an `*Iterate[T]`.
func (i *chainIterator[T, FI, SI, MAP]) Iter() Iterate[T, Iterator[T], MAP] {
	return IterMap[T, Iterator[T], MAP](i)
}
