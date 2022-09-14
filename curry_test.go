package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sum(a, b int) int {
	return a + b
}

func multiply(a, b, c int) int {
	return a * b * c
}

func TestCurry(t *testing.T) {
	t.Run("curry2", func(t *testing.T) {
		add5 := Curry2(sum)(5)

		assert.Equal(t, 15, add5(10))
	})

	t.Run("curry3", func(t *testing.T) {
		multi := Curry3(multiply)

		assert.Equal(t, 24, multi(2)(3)(4))
	})
}
