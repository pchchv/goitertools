package goitertools

import (
	"testing"

	"github.com/pchchv/express/optionext"
	. "github.com/pchchv/go-assert"
	mapext "github.com/pchchv/goitertools/map"
)

func TestMap(t *testing.T) {
	// Test Misc
	iter := WrapMap(makeMap())
	Equal(t, iter.Len(), 5)

	// Test Next
	iter = WrapMap(makeMap())
	Equal(t, iter.Next().IsSome(), true)
	Equal(t, iter.Next().IsSome(), true)
	Equal(t, iter.Next().IsSome(), true)
	Equal(t, iter.Next().IsSome(), true)
	Equal(t, iter.Next().IsSome(), true)
	Equal(t, iter.Next().IsSome(), false)

	// Test Retain
	iter = WrapMap(makeMap()).Retain(func(key string, value int) bool {
		return value == 3
	})
	Equal(t, iter.Next(), optionext.Some(Entry[string, int]{Key: "3", Value: 3}))
	Equal(t, iter.Next(), optionext.None[Entry[string, int]]())

	// Test Iter Filter
	iter2 := WrapMap(makeMap()).Iter().Filter(func(v Entry[string, int]) bool {
		return v.Value != 3
	})
	Equal(t, iter2.Next(), optionext.Some(Entry[string, int]{Key: "3", Value: 3}))
	Equal(t, iter2.Next(), optionext.None[Entry[string, int]]())

	// Test Iter Map
	iterMap := WrapMapWithMap[string, int, int](makeMap()).Iter().Map(func(v Entry[string, int]) int {
		return v.Value
	}).Iter().CollectIter().Sort(func(i int, j int) bool {
		return i < j
	})
	Equal(t, iterMap.Next(), optionext.Some(1))
	Equal(t, iterMap.Next(), optionext.Some(2))
	Equal(t, iterMap.Next(), optionext.Some(3))
	Equal(t, iterMap.Next(), optionext.Some(4))
	Equal(t, iterMap.Next(), optionext.Some(5))
	Equal(t, iterMap.Next(), optionext.None[int]())
}

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
