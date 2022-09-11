package main

func Map[T any, U any](fn func(b T) U, a []T) []U {
	result := make([]U, len(a))

	for i, v := range a {
		result[i] = fn(v)
	}

	return result
}
