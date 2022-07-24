package leetcode

import (
	"fmt"
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

func TestTwoSumCases(t *testing.T) {
	cases := []struct {
		Numbers  []int
		Target   int
		Expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v %d", tc.Numbers, tc.Target), func(t *testing.T) {
			actual := twoSumCache(tc.Numbers, tc.Target)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
