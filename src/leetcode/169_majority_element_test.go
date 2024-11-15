package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1. tally counts
create a map of number => counter
iterate nums

	increment map counter

2. find largest count (most common)
largestKey = 0
largestVal = 0
for num,count := range map:

	if map[num] > largestVal
		largestKey = num
		largestVal = largestVal

return largestVal
*/
func majorityElement(nums []int) int {
	// tally
	counter := map[int]int{}
	for _, num := range nums {
		val, ok := counter[num]
		if ok {
			counter[num] = val + 1
		} else {
			counter[num] = 1
		}
	}

	// find largest
	largestNum := 0
	largestCount := 0
	for num, count := range counter {
		if count > largestCount {
			largestNum = num
			largestCount = count
		}
	}
	return largestNum
}

func Test_169_Example1(t *testing.T) {
	nums := []int{3, 2, 3}
	output := 3
	assert.Equal(t, output, majorityElement(nums))
}

func Test_169_Example2(t *testing.T) {
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
