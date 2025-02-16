package dsa_test

import (
	"reflect"
	"testing"

	"github.com/ericbutera/code-challenges-go/dsa"
)

func TestMaxSlidingWindow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic case",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "single element window",
			nums: []int{1, 3, 1, 2},
			k:    1,
			want: []int{1, 3, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := dsa.MaxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestLongestSubstringKDistinct(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "basic case",
			s:    "eceba",
			k:    2,
			want: 3,
		},
		{
			name: "all unique characters",
			s:    "abcdef",
			k:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := dsa.LongestSubstringKDistinct(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("longestSubstringKDistinct(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

func TestTumblingWindowSum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic case",
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			want: []int{6, 15},
		},
		{
			name: "uneven size",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{3, 7, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := dsa.TumblingWindowSum(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TumblingWindowSum(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
