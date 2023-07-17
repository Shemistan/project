package tests

import (
	"testing"

	"github.com/Edazila/test/lesson1/task1"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("1 + 2", func(t *testing.T) {
		res := task1.RunSum(1, 2)
		expected := 3
		assert.Equal(t, expected, res)
	})

	t.Run("0 + 0", func(t *testing.T) {
		res := task1.RunSum(0, 0)
		expected := 0
		assert.Equal(t, expected, res)
	})

	t.Run("-1 + -1", func(t *testing.T) {
		res := task1.RunSum(-1, -1)
		expected := -2
		assert.Equal(t, expected, res)
	})
}
