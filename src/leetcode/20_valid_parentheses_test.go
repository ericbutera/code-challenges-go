package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.com/problems/valid-parentheses/description/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.


Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.


openers: ( { [
closers: ) } ]

open = []

if current == opener
	open.push(current)

if current == closer
	if !open.last closer for current -> error
	open.pop

scenarios

'()'

loop:
	0:
		current = '('
		open.push(')') push the closer for the current opener
	1:
		current = ')'
		if is_closer
			if open.last != current
				invalid!
			open.pop
		else
			open.push(current)

	at end of loop
	if open.length -> error: there are still opened tags
*/

func isOpener(c string) bool { // todo in array
	// TODO: in array
	// var openers = []string{"(", "{", "["}
	// var closers = []string{")", "}", "]"}
	return c == "(" || c == "{" || c == "["
}

func closerFor(c string) string {
	// TODO: use map
	// var closerFor = []map[string]string{
	// 	"(": ")",
	// 	"[", "]",
	// 	"{", "}",
	// }
	if c == "(" {
		return ")"
	}
	if c == "{" {
		return "}"
	}
	if c == "[" {
		return "]"
	}
	panic("lies")
}

func isValid(s string) bool {
	closers := make([]string, 0) // keeps track of required closers

	for _, c := range s {
		current := string(c) // todo char type?
		if isOpener(current) {
			// append the matching closer
			closer := closerFor(current)
			closers = append(closers, closer)
		} else {
			last := len(closers) - 1

			hasCloser := last < 0
			if hasCloser {
				return false
			}

			unexpectedCloser := current != closers[last]
			if unexpectedCloser {
				return false
			}

			closers = closers[:last] // remove last closer
		}
	}

	return len(closers) == 0
}

func Test_20(t *testing.T) {
	cases := []struct {
		Name   string
		Input  string
		Expect bool
	}{
		{
			Name:   "example 1",
			Input:  "()",
			Expect: true,
		},
		{
			Name:   "example 2",
			Input:  "()[]{}",
			Expect: true,
		},
		{
			Name:   "example 3",
			Input:  "(]",
			Expect: false,
		},
		{
			Name:   "example 4",
			Input:  "([])",
			Expect: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expect, isValid(tc.Input))
		})
	}
}
