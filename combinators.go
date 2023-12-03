package combinators

type Slice[T any] []T

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
