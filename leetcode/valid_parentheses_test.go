package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
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
			actual := isValidParentheses(tc.Input)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
