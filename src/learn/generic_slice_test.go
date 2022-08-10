// https://github.com/ssoroka/slice
package learn

import (
	"testing"

	"github.com/ssoroka/slice"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	input := []int{1, 1, 1, 1}
	result := slice.Unique(input)
	assert.Equal(t, result, []int{1})
}

func TestMap(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := slice.Map[string, string](input, func(i int, s string) string {
		return s + s
	})
	assert.Equal(t, result, []string{"aa", "bb", "cc"})
}
