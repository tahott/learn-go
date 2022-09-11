package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	t.Run("accumulate", func(t *testing.T) {
		target := []int{1, 2, 3}
		res := Fold(func(a, b int) int { return a + b }, 0, target)

		assert.Equal(t, 6, res)
	})

	t.Run("like map function", func(t *testing.T) {
		target := []int{1, 2, 3}
		res := Fold(func(a []int, b int) []int {
			return append(a, b+1)
		}, make([]int, 0), target)

		assert.Equal(t, []int{2, 3, 4}, res)
	})

	t.Run("return only even numbers", func(t *testing.T) {
		target := []int{1, 2, 3, 4, 5}
		res := Fold(func(a []int, b int) []int {
			if b%2 == 0 {
				return append(a, b)
			}

			return a
		}, make([]int, 0), target)

		assert.Equal(t, []int{2, 4}, res)
	})
}
