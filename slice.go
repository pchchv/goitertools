package goitertools

import (
	"github.com/pchchv/express/optionext"
	"github.com/pchchv/goitertools/slice"
)

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

// Map maps a slice of []T -> MAP using the map function.
func (i sliceWrapper[T, MAP]) Map(init MAP, fn func(accum MAP, v T) MAP) MAP {
	return slice.Map[T, MAP](i.slice, init, fn)
}

// Slice returns the underlying sliceWrapper wrapped by the *sliceWrapper.
func (i sliceWrapper[T, MAP]) Slice() []T {
	return i.slice
}

// WrapSliceMap accepts and turns a sliceWrapper into an
// iterator with a map type specified for IterPar() to
// allow the Map helper function.
func WrapSliceMap[T, MAP any](slice []T) sliceWrapper[T, MAP] {
	return sliceWrapper[T, MAP]{
		slice: slice,
	}
}
