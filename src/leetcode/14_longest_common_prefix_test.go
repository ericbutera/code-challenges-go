package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:

    1 <= strs.length <= 200
    0 <= strs[i].length <= 200
    strs[i] consists of only lowercase English letters.

*/

/*
	strs:
	flower
	flow
	flight

	incrementing cursor

	0
	[f]lower
	[f]low
	[f]light

	1
	f[l]ower
	f[l]ow
	f[l]ight

	2
	fl[o]wer
	fl[o]w
	fl[i]ght <- mismatch

	bounds check to ensure current word has enough letters
*/

func longestCommonPrefix(words []string) string {
	wordCount := len(words)
	if wordCount == 0 {
		return ""
	} else if wordCount == 1 {
		return words[0] // prefix is word
	}

	var prefix string

	// note: hasOffset can be removed if we first find the smallest word before iterating

	matchWord := words[0]         // matchWord = flower; use first word as bounds check
	matchLength := len(matchWord) // matchLength = 6; len(flower)
	for cursor := 0; cursor < matchLength; cursor++ {
		matchLetter := matchWord[cursor] // matchLetter = f
		for _, word := range words {     // word = flower
			hasOffset := len(word) > cursor
			if !hasOffset {
				return prefix
			}

			letter := word[cursor] // letter = f
			mismatch := matchLetter != letter
			if mismatch {
				return prefix // mismatch, use longest existing prefix
			}
		}
		// all letters matched cursor, update common prefix
		prefix = matchWord[0 : cursor+1]
	}

	return prefix
}

// func Test_14_Example1(t *testing.T) {
// 	assert.Equal(t, "fl", longestCommonPrefix([]string{
// 		"flower",
// 		"flow",
// 		"flight",
// 	}))
// }
// func Test_14_Example2(t *testing.T) {
// 	assert.Equal(t, "", longestCommonPrefix([]string{
// 		"dog",
// 		"racecar",
// 		"car",
// 	}))
// }

func Test_14(t *testing.T) {
	cases := []struct {
		Name     string
		Input    []string
		Expected string
	}{
		{
			Name:     "no words",
			Input:    []string{},
			Expected: "",
		},
		{
			Name:     "one word",
			Input:    []string{"one"},
			Expected: "one",
		},
		{
			Name:     "example 1",
			Input:    []string{"flower", "flow", "flight"},
			Expected: "fl",
		},
		{
			Name:     "example 2",
			Input:    []string{"dog", "racecar", "car"},
			Expected: "",
		},
		{
			Name:     "matching words",
			Input:    []string{"flower", "flower"},
			Expected: "flower",
		},
		{
			Name:     "longer first word",
			Input:    []string{"ab", "a"},
			Expected: "a",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, longestCommonPrefix(tc.Input))
		})
	}
}
