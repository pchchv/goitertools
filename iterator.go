package goitertools

import (
	"runtime"
	"sync"

	"github.com/pchchv/express/optionext"
)

var numCPU = runtime.NumCPU()

// Iterator is an interface that represents
// something performing an iteration using the Next method.
type Iterator[T any] interface {
	// Next advances the iterator and returns the next value.
	//
	// Returns an Option with value of None when iteration has finished.
	Next() optionext.Option[T]
}

// PeekableIterator is an interface representing something that
// iterates using the Next method and ability to `Peek` the
// next element value without advancing the `Iterator`.
type PeekableIterator[T any] interface {
	Iterator[T]
	// Peek returns the `Next` element from the `Iterator` without advancing it.
	Peek() optionext.Option[T]
}

// Iterate is an iterator with attached helper functions.
type Iterate[T any, I Iterator[T], MAP any] struct {
	iterator I
}

// Next returns the new iterator value.
func (i Iterate[T, I, MAP]) Next() optionext.Option[T] {
	return i.iterator.Next()
}

// forEach is an early cancellable form of `ForEach`.
func (i Iterate[T, I, MAP]) forEach(parallel bool, fn func(T) (stop bool)) {
	if parallel {
		var stopOnce sync.Once
		stopEarly := make(chan struct{})
		in := make(chan optionext.Option[T])
		wg := new(sync.WaitGroup)
		for j := 0; j < numCPU; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
			FOR:
				for {
					select {
					case <-stopEarly:
						break FOR
					case v := <-in:
						if v.IsNone() || fn(v.Unwrap()) {
							stopOnce.Do(func() {
								close(stopEarly)
							})
							break FOR
						}

					}
				}
			}()
		}
	FOR:
		for {
			select {
			case <-stopEarly:
				break FOR
			case in <- i.iterator.Next():
			}
		}
		close(in)
		wg.Wait()
	} else {
		for {
			v := i.iterator.Next()
			if v.IsNone() || fn(v.Unwrap()) {
				break
			}
		}
	}
}
