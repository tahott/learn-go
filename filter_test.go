package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("filter even number", func(t *testing.T) {
		target := []int{1, 2, 3, 4, 5}
		res := Filter(func(x int) bool { return x%2 == 0 }, target)

		assert.Equal(t, []int{2, 4}, res)
	})
}
