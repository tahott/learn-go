package main

import "math"

func Take[T any](cnt uint, src []T) []T {
	count := int(math.Min(float64(cnt), float64(len(src))))
	result := make([]T, 0)

	result = append(result, src[:count]...)

	return result
}
