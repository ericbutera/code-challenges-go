package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findIntersectionValues_BruteForce(nums1 []int, nums2 []int) []int {
	// brute force comparison
	// space efficient, runtime inefficient
	// O(n^2)
	len1 := len(nums1)
	len2 := len(nums2)

	answer1 := 0
	answer2 := 0
	for x := 0; x < len1; x++ {
		for y := 0; y < len2; y++ {
			if nums1[x] == nums2[y] {
				answer1++
				break
			}
		}
	}
	for y := 0; y < len2; y++ {
		for x := 0; x < len1; x++ {
			if nums2[y] == nums1[x] {
				answer2++
				break
			}
		}
	}

	return []int{
		answer1,
		answer2,
	}
}

func findIntersectionValues_Map(nums1 []int, nums2 []int) []int {
	// build map for constant time check if num exists
	// space inefficient, runtime efficient
	set1 := map[int]bool{}
	for _, num := range nums1 {
		set1[num] = true
	}

	set2 := map[int]bool{}
	for _, num := range nums2 {
		set2[num] = true
	}

	answer1 := 0
	answer2 := 0
	for _, num := range nums1 {
		if _, ok := set2[num]; ok {
			fmt.Printf("1 num %v\n", num)
			answer1++
		}
	}

	for _, num := range nums2 {
		if _, ok := set1[num]; ok {
			fmt.Printf("2 num %v\n", num)
			answer2++
		}
	}

	return []int{answer1, answer2}
}

func Test_2956(t *testing.T) {
	cases := []struct {
		Name   string
		InputA []int
		InputB []int
		Expect []int
	}{
		{
			Name:   "example 1",
			InputA: []int{2, 3, 2},
			InputB: []int{1, 2},
			Expect: []int{2, 1},
		},

		{
			Name:   "case 2",
			InputA: []int{4, 3, 2, 3, 1},
			InputB: []int{2, 2, 5, 2, 3, 6},
			Expect: []int{3, 4},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expect, findIntersectionValues_BruteForce(tc.InputA, tc.InputB))
			assert.Equal(t, tc.Expect, findIntersectionValues_Map(tc.InputA, tc.InputB))
		})
	}
}
