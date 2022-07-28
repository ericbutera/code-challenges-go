package learn

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	actual := strings.Contains("seafood", "oo")
	assert.Equal(t, actual, true)
}

func TestChar(t *testing.T) {
	word := "Hello"
	assert.Equal(t, "e", string(word[1]))
}

func TestSubString(t *testing.T) {
	tale := "we were all going direct to Heaven, we were all going direct the other way"
	assert.Equal(t, "we", string(tale[:2]))
	assert.Equal(t, "other way", string(tale[65:]))
}
