package main

func Filter[T any](fn func(b T) bool, a []T) []T {
	result := make([]T, 0)

	for _, v := range a {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}
