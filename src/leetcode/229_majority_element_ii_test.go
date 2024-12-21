package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func majorityElement(nums []int) []int {
	counter := make(map[int]int) // frequency counter

	for _, num := range nums {
		counter[num]++
	}

	threshold := len(nums) / 3

	majority := []int{}
	for num, count := range counter {
		if count > threshold {
			majority = append(majority, num)
		}
	}

	return majority
}

func Test_229(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Nums     []int
		Expected []int
	}{
		{
			Nums:     []int{3, 2, 3},
			Expected: []int{3},
		},
		{
			Nums:     []int{1},
			Expected: []int{1},
		},
		{
			Nums:     []int{1, 2},
			Expected: []int{1, 2},
		},
		{
			Nums:     []int{1, 1, 2, 2, 1, 1, 1, 1, 2, 2, 1, 1},
			Expected: []int{1},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.Nums), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, majorityElement(tc.Nums))
		})
	}
}
