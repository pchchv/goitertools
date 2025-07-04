package slice

import (
	"strconv"
	"testing"

	"github.com/pchchv/express/optionext"
	. "github.com/pchchv/go-assert"
)

func TestSort(t *testing.T) {
	s := []int{0, 1, 2}
	Sort(s, func(i int, j int) bool {
		return i > j
	})
	Equal(t, s[0], 2)
	Equal(t, s[1], 1)
	Equal(t, s[2], 0)
}

func TestSortStable(t *testing.T) {
	s := []int{0, 1, 1, 2}
	SortStable(s, func(i int, j int) bool {
		return i > j
	})
	Equal(t, s[0], 2)
	Equal(t, s[1], 1)
	Equal(t, s[2], 1)
	Equal(t, s[3], 0)
}

func TestReverse(t *testing.T) {
	s := []int{1, 2}
	Reverse(s)
	Equal(t, []int{2, 1}, s)

	s = []int{1, 2, 3}
	Reverse(s)
	Equal(t, []int{3, 2, 1}, s)
}

func TestRetain(t *testing.T) {
	s := Retain([]int{0, 1, 2, 3}, func(v int) bool {
		return v > 0 && v < 3
	})
	Equal(t, len(s), 2)
	Equal(t, s[0], 1)
	Equal(t, s[1], 2)
}

func TestReduce(t *testing.T) {
	result := Reduce([]int{0, 1, 2}, func(accum int, current int) int {
		return accum + current
	})
	Equal(t, result, optionext.Some(3))

	// Test Reduce empty slice
	result = Reduce([]int{}, func(accum int, current int) int {
		return accum + current
	})
	Equal(t, result, optionext.None[int]())
}

func TestFilter(t *testing.T) {
	s := Filter([]int{0, 1, 2, 3}, func(v int) bool {
		return v > 0 && v < 3
	})
	Equal(t, len(s), 2)
	Equal(t, s[0], 0)
	Equal(t, s[1], 3)

}

func TestMap(t *testing.T) {
	s := Map([]int{0, 1, 2, 3}, make([]string, 0, 4), func(accum []string, v int) []string {
		return append(accum, strconv.Itoa(v))
	})
	Equal(t, len(s), 4)
	Equal(t, s[0], "0")
	Equal(t, s[1], "1")
	Equal(t, s[2], "2")
	Equal(t, s[3], "3")

	s2 := Map(nil, nil, func(accum []string, v int) []string {
		return append(accum, strconv.Itoa(v))
	})
	Equal(t, len(s2), 0)
}
