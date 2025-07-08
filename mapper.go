package goitertools

// MapFn represents the mapWrapper transformation function.
type MapFn[T, MAP any] func(v T) MAP

// mapWrapper is used to transform elements from one type to another.
type mapper[T any, I Iterator[T], MAP any] struct {
	iterator I
	fn       MapFn[T, MAP]
}
