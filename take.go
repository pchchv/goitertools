package goitertools

// takeIterator is an iterator that only iterates over n elements.
type takeIterator[T any, I Iterator[T], MAP any] struct {
	iterator I
	limit    int
}

// Next returns the next element until n is reached or end of the iterator.
func (i *takeIterator[T, I, MAP]) Next() optionext.Option[T] {
	if i.limit <= 0 {
		return optionext.None[T]()
	}
	i.limit--
	return i.iterator.Next()
}

// Iter is a convenience function that converts the `takeIterator` iterator into an `*Iterate[T]`.
func (i *takeIterator[T, I, MAP]) Iter() Iterate[T, Iterator[T], MAP] {
	return IterMap[T, Iterator[T], MAP](i)
}
