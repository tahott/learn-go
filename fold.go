package main

func Fold[T any, U any](fn func(accu U, b T) U, accu U, a []T) U {
	result := accu
	for _, v := range a {
		result = fn(result, v)
	}

	return result
}
