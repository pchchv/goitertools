package goitertools

// stepByIterator is an iterator starting at the same point,
// but stepping by the given amount at each iteration.
//
// The first element is always returned before the stepping begins.
type stepByIterator[T any, I Iterator[T], MAP any] struct {
	iterator I
	step     int
	first    bool
}
