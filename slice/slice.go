package slice

import "sort"

// Sort sorts the sliceWrapper x given the provided less function.
//
// The sort is not guaranteed to be stable:
// equal elements may be reversed from their original order.
//
// For a stable sort, use SortStable.
func Sort[T any](slice []T, less func(i T, j T) bool) {
	sort.Slice(slice, func(j, k int) bool {
		return less(slice[j], slice[k])
	})
}

// SortStable sorts the sliceWrapper x using the provided less function,
// keeping equal elements in their original order.
func SortStable[T any](slice []T, less func(i T, j T) bool) {
	sort.SliceStable(slice, func(j, k int) bool {
		return less(slice[j], slice[k])
	})
}
