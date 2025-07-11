package goitertools

// chunker chunks the returned elements into slices of specified size.
//
// The last returned slice is NOT guaranteed to be
// of exact size unless the input exactly lines up.
type chunker[T any, I Iterator[T], MAP any] struct {
	iterator I
	size     int
}
