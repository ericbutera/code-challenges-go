package leetcode_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inversed = map[string]string{ //nolint:gochecknoglobals
	"}": "{",
	"]": "[",
	")": "(",
}

const (
	Starts = "[{("
	Ends   = "]})"
)

// https://www.educative.io/answers/how-to-implement-a-stack-in-golang
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func isValidParentheses(input string) bool {
	var stack Stack

	for i := 0; i < len(input); i++ { //nolint:intrange
		char := string(input[i])
		if !parseChar(char, &stack) {
			return false
		}
	}

	return len(stack) == 0
}

func parseChar(char string, stack *Stack) bool {
	if strings.Contains(Starts, char) {
		stack.Push(char)
		return true
	}

	if strings.Contains(Ends, char) {
		if len(*stack) == 0 {
			return false
		}

		last, exists := stack.Pop()
		if exists {
			inverse := inversed[char]
			if last == inverse {
				return true
			}
		}
	}

	return false
}

func TestValidParentheses(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Input    string
		Expected bool
	}{
		{"()", true},
		{"(]", false},
		{"([", false},
		{"{[]}", true},
		{"(({}))", true},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v %v", tc.Input, tc.Expected), func(t *testing.T) {
			t.Parallel()
			actual := isValidParentheses(tc.Input)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
