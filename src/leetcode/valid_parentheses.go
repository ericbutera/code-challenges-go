package leetcode

import "strings"

var Inverse = map[string]string{
	"}": "{",
	"]": "[",
	")": "(",
}

const Starts = "[{("
const Ends = "]})"

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
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func isValidParentheses(input string) bool {
	var stack Stack

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if !parseChar(char, &stack) {
			return false
		}
	}

	return len(stack) <= 0
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
			inverse := Inverse[char]
			if last == inverse {
				return true
			}
		}
	}

	return false
}
