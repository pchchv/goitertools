package goitertools

import (
	"testing"

	mapext "github.com/pchchv/goitertools/map"
)

func BenchmarkRetainMap_Retain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapext.Retain(makeMap(), func(_ string, value int) (retain bool) {
			return value == 3
		})
	}
}

func makeMap() map[string]int {
	return map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}
}
