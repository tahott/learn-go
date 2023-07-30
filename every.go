package main

func Every[T any](fn func(b T) bool, a []T) bool {
	for _, v := range a {
		if !fn(v) {
			return false
		}
	}

	return true
}
