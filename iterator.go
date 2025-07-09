package goitertools

import (
	"runtime"
	"sync"
	"sync/atomic"

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

// Any returns true if any element matches the function return,
// false otherwise.
func (i Iterate[T, I, MAP]) Any(fn func(T) bool) (isAny bool) {
	i.forEach(false, func(v T) (stop bool) {
		isAny = fn(v)
		return isAny
	})
	return
}

// AnyParallel returns true if any element matches the function return, false otherwise.
//
// This will run in parallel.
// It is recommended to only use this when the
// overhead of running n parallel is less than the work needing to be done.
func (i Iterate[T, I, MAP]) AnyParallel(fn func(T) bool) (isAny bool) {
	var k uint32
	i.forEach(true, func(v T) (stop bool) {
		match := fn(v)
		if match {
			atomic.StoreUint32(&k, 1)
		}
		return match
	})
	return k == 1
}

// All returns true if all element matches the function return,
// false otherwise.
func (i Iterate[T, I, MAP]) All(fn func(T) bool) (isAll bool) {
	var checked bool
	i.forEach(false, func(v T) (stop bool) {
		checked = fn(v)
		return !checked
	})
	return checked
}

// AllParallel returns true if all element matches the function return, false otherwise.
//
// This will run in parallel.
// It is recommended to only use this when the
// overhead of running n parallel is less than the work needing to be done.
func (i Iterate[T, I, MAP]) AllParallel(fn func(T) bool) (isAll bool) {
	var k uint32 = 1
	i.forEach(true, func(v T) (stop bool) {
		if fn(v) {
			return false
		}

		atomic.StoreUint32(&k, 0)
		return true
	})
	return k == 1
}

// ForEach runs the provided function for each element until completion.
//
// This will run in parallel is using a parallel iterator.
func (i Iterate[T, I, MAP]) ForEach(fn func(T)) {
	i.forEach(false, func(t T) (stop bool) {
		fn(t)
		return false
	})
}

// ForEachParallel runs the provided function for each element in parallel until completion.
//
// The function must maintain its own thread safety.
func (i Iterate[T, I, MAP]) ForEachParallel(fn func(T)) {
	i.forEach(true, func(t T) (stop bool) {
		fn(t)
		return false
	})
}

// Count consumes the iterator and returns count if iterations.
//
// This will run in parallel is using a parallel iterator.
func (i Iterate[T, I, MAP]) Count() (j int) {
	i.ForEach(func(_ T) {
		j++
	})
	return j
}

// CountParallel consumes the iterator concurrently and returns count if iterations.
func (i Iterate[T, I, MAP]) CountParallel() int {
	var j int64
	i.ForEach(func(_ T) {
		atomic.AddInt64(&j, 1)
	})
	return int(j)
}

// Position searches for an element in an iterator, returning its index.
func (i Iterate[T, I, MAP]) Position(fn func(T) bool) optionext.Option[int] {
	for j := 0; ; j++ {
		result := i.iterator.Next()
		if result.IsNone() {
			return optionext.None[int]()
		} else if fn(result.Unwrap()) {
			return optionext.Some(j)
		}
	}
}

// Partition creates two collections from supplied function,
// all elements returning true will be in one result and all that
// were returned false in the other.
func (i Iterate[T, I, MAP]) Partition(fn func(v T) bool) (left, right []T) {
	i.ForEach(func(v T) {
		if fn(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	})
	return
}

// Find searches for the next element of an iterator that satisfies the function.
func (i Iterate[T, I, MAP]) Find(fn func(T) bool) (result optionext.Option[T]) {
	for {
		result = i.iterator.Next()
		if result.IsNone() || fn(result.Unwrap()) {
			return
		}
	}
}

// Collect transforms an iterator into a sliceWrapper.
//
// This will run in parallel is using a parallel iterator.
func (i Iterate[T, I, MAP]) Collect() (results []T) {
	i.ForEach(func(v T) {
		results = append(results, v)
	})
	return
}

// Reduce reduces the elements to a single one,
// by repeatedly applying a reducing function.
func (i Iterate[T, I, MAP]) Reduce(fn func(accum T, current T) T) optionext.Option[T] {
	v := i.iterator.Next()
	if v.IsNone() {
		return optionext.None[T]()
	}

	accum := v.Unwrap()
	for {
		current := i.iterator.Next()
		if current.IsNone() {
			return optionext.Some(accum)
		}
		accum = fn(accum, current.Unwrap())
	}
}

// Map accepts a `FilterFn[T]` to filter items.
//
// NOTE: This is made possible by passing the one-time possible MAP type around. This is unfortunate but the only way it
// can be supported due to the limitations of the Go Compiler.
func (i Iterate[T, I, MAP]) Map(fn MapFn[T, MAP]) mapper[T, Iterator[T], MAP] {
	return Map[T, Iterator[T], MAP](i.iterator, fn)
}

// Filter accepts a `FilterFn[T]` to filter items.
func (i Iterate[T, I, MAP]) Filter(fn FilterFn[T]) Iterate[T, Iterator[T], MAP] {
	return IterMap[T, Iterator[T], MAP](FilterWithMap[T, I, MAP](i.iterator, fn))
}

// CollectIter transforms an iterator into a sliceWrapper and
// returns a *sliceWrapper in order to run additional functions inline such as Sort().
//
// eg. .Filter(...).CollectIter().Sort(...).WrapSlice()
//
// This will run in parallel is using a parallel iterator.
func (i Iterate[T, I, MAP]) CollectIter() sliceWrapper[T, MAP] {
	return WrapSliceMap[T, MAP](i.Collect())
}

// PartitionIter creates two iterable collections from supplied function,
// all elements returning true will be in one result
// and all that were returned false in the other.
func (i Iterate[T, I, MAP]) PartitionIter(fn func(v T) bool) (left, right sliceWrapper[T, struct{}]) {
	l, r := i.Partition(fn)
	return WrapSlice(l), WrapSlice(r)
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

// Iter creates a new iterator with helper functions.
//
// It defaults the Map() function to struct{}.
// Use IterMap() if you wish to specify a type.
func Iter[T any, I Iterator[T]](iterator I) Iterate[T, I, struct{}] {
	return IterMap[T, I, struct{}](iterator)
}

// IterMap creates a new iterator with helper functions.
//
// It accepts a map type `MAP` to allow for usage of the `Map` and `CollectMap` helper function inline.
// You must use the Map() function standalone otherwise.
func IterMap[T any, I Iterator[T], MAP any](iterator I) Iterate[T, I, MAP] {
	return Iterate[T, I, MAP]{
		iterator: iterator,
	}
}
