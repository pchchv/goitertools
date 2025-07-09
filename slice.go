package goitertools

import "github.com/pchchv/express/optionext"

type sliceWrapper[T, MAP any] struct {
	slice []T
}

// Len returns the length of the underlying sliceWrapper.
func (i sliceWrapper[T, MAP]) Len() int {
	return len(i.slice)
}

func (i *sliceWrapper[T, MAP]) Next() optionext.Option[T] {
	if len(i.slice) == 0 {
		return optionext.None[T]()
	}

	v := i.slice[0]
	i.slice = i.slice[1:]
	return optionext.Some(v)
}

// IntoIter turns the slice wrapper into an `Iterator[T]`
func (i sliceWrapper[T, MAP]) IntoIter() *sliceWrapper[T, MAP] {
	return &i
}

// Iter is a convenience function that converts the sliceWrapper iterator into an `*Iterate[T]`.
func (i sliceWrapper[T, MAP]) Iter() Iterate[T, *sliceWrapper[T, MAP], MAP] {
	return IterMap[T, *sliceWrapper[T, MAP], MAP](i.IntoIter())
}
