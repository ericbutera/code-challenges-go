package anki_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func compress(in string) string {
	out := ""

	len := len(in)
	if len == 0 || len == 1 {
		return in
	}

	previous := string(in[0]) // prevent mismatch on first char
	count := 0
	for _, c := range in {
		current := string(c)
		if current != previous {
			out += format(previous, count)
			count = 1
			previous = current
		} else {
			count += 1
		}
	}

	out += format(previous, count)

	return out
}

func format(previous string, count int) string {
	if count > 1 {
		return fmt.Sprintf("%s%d", previous, count)
	}
	return previous
}

func Test_Compress_Example1(t *testing.T) {
	assert.Equal(t, "A3BC2D4", compress("AAABCCDDDD"))
}

func Test_Compress_Example2(t *testing.T) {
	assert.Equal(t, "ABC", compress("ABC"))
}

func Test_Compress_Example3(t *testing.T) {
	assert.Equal(t, "ABC2", compress("ABCC"))
}
