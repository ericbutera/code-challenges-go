package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strStr(haystack string, needle string) int {
	// outer loop haystack
	// - check haystack[x]
	// - inner loop needle[y]
	//   - iter 0 compare haystack[x0] == needle[y0]
	//   - iter 1 compare haystack[x0+y1] == needle[y1]
	//   - iter 2 compare haystack[x0+y2] == needle[y2]
	if needle == haystack {
		return 0
	}
	haystackLength := len(haystack)
	needleLength := len(needle)

	if needleLength > haystackLength {
		return -1
	}

	for x := 0; x < haystackLength; x++ {
		match := true
		for y := 0; y < needleLength; y++ {
			// bounds check next & y
			next := x + y
			if next > haystackLength-1 { // invalid index
				match = false
				break
			}
			if y > needleLength-1 {
				match = false
				break
			}

			mismatch := haystack[next] != needle[y]
			if mismatch {
				match = false
				break
			}
		}
		if match {
			return x
		}
	}
	return -1
}

func Test_28(t *testing.T) {
	cases := []struct {
		Name     string
		Haystack string
		Needle   string
		Expected int
	}{
		{
			Name:     "example 1",
			Haystack: "sadbutsad",
			Needle:   "sad",
			Expected: 0,
		},
		{
			Name:     "example 2",
			Haystack: "leetcode",
			Needle:   "leeto",
			Expected: -1,
		},
		{
			Name:     "example 01",
			Haystack: "sadbutsab",
			Needle:   "sab",
			Expected: 6,
		},
		{
			Name:     "example 01",
			Haystack: "happy",
			Needle:   "habbyhappy",
			Expected: -1,
		},
		{
			Name:     "example 02",
			Haystack: "abc",
			Needle:   "c",
			Expected: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := strStr(tc.Haystack, tc.Needle)
			assert.Equal(t, tc.Expected, actual, "needle:%s haystack:%s", tc.Needle, tc.Haystack)
		})
	}
}
