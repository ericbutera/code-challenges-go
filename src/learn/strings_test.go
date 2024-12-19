package learn_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()
	actual := strings.Contains("seafood", "oo")
	assert.True(t, actual)
}

func TestChar(t *testing.T) {
	t.Parallel()
	word := "Hello"
	assert.Equal(t, "e", string(word[1]))
}

func TestSubString(t *testing.T) {
	t.Parallel()
	tale := "we were all going direct to Heaven, we were all going direct the other way"
	assert.Equal(t, "we", tale[:2])
	assert.Equal(t, "other way", tale[65:])
}
