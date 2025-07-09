package goitertools

import (
	"testing"

	"github.com/pchchv/goitertools/slice"
)

func BenchmarkSTDRetain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdRetain(makeSlice())
	}
}

func BenchmarkSTDFnRetain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdRetainFn(makeSlice(), func(v int) bool {
			return v == 1
		})
	}
}

func BenchmarkSliceWrapper_Retain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WrapSlice(makeSlice()).Retain(func(v int) bool {
			return v == 1
		})
	}
}

func BenchmarkRetainSlice_Retain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice.Retain(makeSlice(), func(v int) bool {
			return v == 1
		})
	}
}

func stdRetain(s []int) []int {
	var j int
	for _, v := range s {
		if v == 1 {
			s[j] = v
			j++
		}
	}

	return s[:j]
}

func stdRetainFn(s []int, fn func(v int) bool) []int {
	var j int
	for _, v := range s {
		if fn(v) {
			s[j] = v
			j++
		}
	}

	return s[:j]
}
