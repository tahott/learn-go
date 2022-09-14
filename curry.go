package main

func Curry2[T any](fn func(a, b T) T) func(a T) func(b T) T {
	return func(a T) func(b T) T {
		return func(b T) T {
			return fn(a, b)
		}
	}
}

func Curry3[T any](fn func(a, b, c T) T) func(a T) func(b T) func(c T) T {
	return func(a T) func(b T) func(c T) T {
		return func(b T) func(c T) T {
			return func(c T) T {
				return fn(a, b, c)
			}
		}
	}
}
