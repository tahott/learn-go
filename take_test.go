package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	t.Run("", func(t *testing.T) {
		target := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		res := Take(5, target)

		assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
	})

	t.Run("", func(t *testing.T) {
		target := []int{1, 2, 3}
		res := Take(5, target)

		assert.Equal(t, []int{1, 2, 3}, res)
	})
}
