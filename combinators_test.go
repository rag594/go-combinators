package combinators

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	intSlice := Slice[int]{1, 2, 3, 4, 5, 6}

	evenPredicateFunc := func(n int) bool {
		return n%2 == 0
	}

	opEvenSlice := Slice[int]{2, 4, 6}

	evenSlice := intSlice.Filter(evenPredicateFunc)
	assert.Equal(t, true, slices.Equal[Slice[int]](evenSlice, opEvenSlice))

}

func TestMap(t *testing.T) {
	intSlice := Slice[int]{1, 2, 3, 4, 5, 6}

	mapperFunc := func(n int) int {
		return n * 2
	}

	multiplierSlice := Slice[int]{2, 4, 6, 8, 10, 12}

	mappedSlice := intSlice.Map(mapperFunc)
	assert.Equal(t, true, slices.Equal[Slice[int]](mappedSlice, multiplierSlice))

}

func TestMapAndFilterCombination(t *testing.T) {
	s := Slice[string]{"abc", "def", "fjkl", "deq", "dqdq"}
	d := s.
		Filter(func(t string) bool { return strings.Contains(t, "e") }).
		Map(func(t string) string {
			return strings.ToUpper(t)
		})
	expectedOutput := Slice[string]{"DEF", "DEQ"}
	assert.Equal(t, true, slices.Equal[Slice[string]](d, expectedOutput))

}
