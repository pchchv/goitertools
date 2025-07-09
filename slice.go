package goitertools

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
