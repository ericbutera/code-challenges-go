package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	return twoSumBruteForce(nums, target)
}

func twoSumBruteForce(nums []int, target int) []int {
	for x := 0; x < len(nums)-1; x++ { //nolint:intrange
		for y := x + 1; y < len(nums); y++ {
			attempt := nums[x] + nums[y]
			if attempt == target {
				return []int{x, y}
			}
		}
	}

	return []int{0, 0}
}

func twoSumCache(nums []int, target int) []int {
	cache := map[int]int{}
	for x, number := range nums {
		search := target - number
		if index, ok := cache[search]; ok {
			return []int{index, x}
		}
		cache[number] = x
	}
	return []int{0, 0}
}

func TestTwoSum(t *testing.T) {
	t.Parallel()
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	// actual := twoSum(nums, target)
	actual := twoSumCache(nums, target)
	assert.Equal(t, expected, actual)
}

func TestTwoSumCases(t *testing.T) {
	t.Parallel()
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
			t.Parallel()
			assert.Equal(t, tc.Expected, twoSumCache(tc.Numbers, tc.Target))
			assert.Equal(t, tc.Expected, twoSum(tc.Numbers, tc.Target))
		})
	}
}
