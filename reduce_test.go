package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Run("element sum", func(t *testing.T) {
		target := []int{1, 2, 3}
		res := Reduce(func(a, b int) int { return a + b }, target)

		assert.Equal(t, 6, res)
	})

	t.Run("string concatenation", func(t *testing.T) {
		target := []string{"hello", "my", "name", "is", "..."}
		res := Reduce(func(a, b string) string { return a + " " + b }, target)

		assert.Equal(t, "hello my name is ...", res)
	})

	t.Run("find max value", func(t *testing.T) {
		target := []int{1, 13, 5, 9, 7, 2}
		res := Reduce(func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}, target)

		assert.Equal(t, 13, res)
	})
}
