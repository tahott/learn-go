package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	name string
	age  uint
}

func TestMap(t *testing.T) {
	t.Run("when add 1 then return", func(t *testing.T) {
		target := []int{1, 2, 3}
		res := Map(func(x int) int { return x + 1 }, target)

		assert.Equal(t, []int{2, 3, 4}, res)
	})

	t.Run("when add prefix 'hello' then return", func(t *testing.T) {
		target := []string{"jimmy", "alice", "hong"}
		res := Map(func(s string) string { return "hello " + s }, target)

		assert.Equal(t, []string{"hello jimmy", "hello alice", "hello hong"}, res)
	})

	t.Run("structure to string", func(t *testing.T) {
		target := []User{{name: "hong", age: 17}, {name: "kim", age: 20}}
		res := Map(func(u User) string { return u.name }, target)

		assert.Equal(t, []string{"hong", "kim"}, res)
	})
}
