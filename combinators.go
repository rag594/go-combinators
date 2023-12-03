package combinators

// Slice Container parametrised by Type T
type Slice[T any] []T

// Filter is a combinator which can be used in conjuction with other
// combinators.
// It is used to Filter out the elements based on the filterFunc
// provided.
// For example:
//
// intSlice := Slice[int]{1, 2, 3, 4, 5, 6}
//
//	evenPredicateFunc := func(n int) bool {
//		return n%2 == 0
//	}
//
// opEvenSlice := Slice[int]{2, 4, 6}
// evenSlice := intSlice.Filter(evenPredicateFunc)
// Output: Slice[int]{2, 4, 6}
func (s Slice[T]) Filter(filterFunc func(T) bool) Slice[T] {
	result := Slice[T]{}

	for _, item := range s {
		r := filterFunc(item)
		if r {
			result = append(result, item)
		}
	}

	return result
}

func (s Slice[T]) Map(mapFunc func(T) T) Slice[T] {
	result := Slice[T]{}

	for _, item := range s {
		result = append(result, mapFunc(item))
	}

	return result
}

func (s Slice[T]) Reduce(startVal *T, callbackFunc func(T, T) T) T {
	var accumulator T
	start := 0
	if startVal != nil {
		accumulator = *startVal
	} else {
		accumulator = s[0]
		start = 1
	}

	for idx := start; idx < len(s); idx++ {
		accumulator = callbackFunc(accumulator, s[idx])
	}

	return accumulator
}
