package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Grade struct {
	Subject string
	Score   uint
}

func TestEvery(t *testing.T) {
	t.Run("Element is every even number", func(t *testing.T) {
		target := []int{2, 4, 6, 8, 10}
		res := Every(func(a int) bool {
			return a%2 == 0
		}, target)

		assert.Equal(t, true, res)
	})

	t.Run("All subject grade are gratest than 60", func(t *testing.T) {
		target := []Grade{
			{Subject: "Korean", Score: 80},
			{Subject: "Math", Score: 90},
			{Subject: "Physics", Score: 60},
			{Subject: "Chemistry", Score: 70},
		}

		res := Every(func(a Grade) bool {
			return a.Score >= 60
		}, target)

		assert.Equal(t, true, res)
	})
}
