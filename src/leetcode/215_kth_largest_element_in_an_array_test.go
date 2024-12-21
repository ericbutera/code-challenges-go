package leetcode_test

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// -10^4 <= nums[i] <= 104
// 1 <= k <= nums.length <= 105
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return nums[k-1]
}

func Test_215(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Nums     []int
		K        int
		Expected int
	}{
		{
			Nums:     []int{3, 2, 1, 5, 6, 4},
			K:        5,
			Expected: 2,
		},
		{
			Nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			K:        4,
			Expected: 4,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v %d", tc.Nums, tc.K), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, findKthLargest(tc.Nums, tc.K))
		})
	}
}
