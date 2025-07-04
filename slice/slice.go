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

// Reverse reverses the slice contents.
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Retain retains only the elements specified by the function.
//
// This returns a new slice with references to the
// underlying data instead of shuffling.
func Retain[T any](slice []T, fn func(v T) bool) []T {
	results := make([]T, 0, len(slice))
	for _, v := range slice {
		v := v
		if fn(v) {
			results = append(results, v)
		}
	}
	return results
}
