package mapext

import (
	"sort"
	"testing"

	. "github.com/pchchv/go-assert"
)

func TestMap(t *testing.T) {
	// Test Map to slice
	m := map[string]int{
		"0": 0,
		"1": 1,
	}
	slice := Map(m, make([]int, 0, len(m)), func(accum []int, key string, value int) []int {
		return append(accum, value)
	})
	sort.SliceStable(slice, func(i, j int) bool {
		return i < j
	})
	Equal(t, len(slice), 2)

	// Test Map to Map of different type
	inverted := Map(m, make(map[int]string, len(m)), func(accum map[int]string, key string, value int) map[int]string {
		accum[value] = key
		return accum
	})
	Equal(t, len(inverted), 2)
	Equal(t, inverted[0], "0")
	Equal(t, inverted[1], "1")
}
