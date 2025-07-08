package goitertools

// MapFn represents the mapWrapper transformation function.
type MapFn[T, MAP any] func(v T) MAP
