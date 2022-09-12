package main

func Reduce[T any](fn func(accu T, b T) T, a []T) T {
	result := a[0]
	for _, v := range a[1:] {
		result = fn(result, v)
	}

	return result
}
