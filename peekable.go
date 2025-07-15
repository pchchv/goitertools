package goitertools

import "github.com/pchchv/express/optionext"

// peekableIterator makes an `Iterator` peekable.
type peekableIterator[T any, I Iterator[T]] struct {
	iterator I
	prev     optionext.Option[T]
}
