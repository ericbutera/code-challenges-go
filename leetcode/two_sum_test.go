package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	// actual := twoSum(nums, target)
	actual := twoSumCache(nums, target)
	assert.Equal(t, expected, actual)
}
