package goitertools

import "github.com/pchchv/express/optionext"

// chunker chunks the returned elements into slices of specified size.
//
// The last returned slice is NOT guaranteed to be
// of exact size unless the input exactly lines up.
type chunker[T any, I Iterator[T], MAP any] struct {
	iterator I
	size     int
}

// Next yields the next set of elements from the iterator.
func (i chunker[T, I, MAP]) Next() optionext.Option[[]T] {
	chunk := make([]T, 0, i.size)
	for {
		v := i.iterator.Next()
		if v.IsNone() {
			break
		}
		chunk = append(chunk, v.Unwrap())
		if len(chunk) == cap(chunk) {
			break
		}
	}

	if len(chunk) == 0 {
		return optionext.None[[]T]()
	}

	return optionext.Some(chunk)
}
