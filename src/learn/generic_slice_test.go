// https://github.com/ssoroka/slice
package learn_test

import (
	"testing"

	"github.com/ssoroka/slice"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	t.Parallel()
	input := []int{1, 1, 1, 1}
	result := slice.Unique(input)
	assert.Equal(t, []int{1}, result)
}

func TestMap(t *testing.T) {
	t.Parallel()
	input := []string{"a", "b", "c"}
	result := slice.Map[string, string](input, func(_ int, s string) string {
		return s + s
	})
	assert.Equal(t, []string{"aa", "bb", "cc"}, result)
}
