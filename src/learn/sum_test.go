package learn_test

import (
	"fmt"
	"testing"

	"github.com/ericbutera/code-challenges-go/src/learn"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()
	numbers := [5]int{1, 2, 3, 4, 5}
	got := learn.Sum(numbers)
	assert.Equal(t, 15, got)
}

func TestAdder(t *testing.T) {
	t.Parallel()
	sum := learn.Add(2, 2)
	assert.Equal(t, 4, sum)
}

// https://go.dev/blog/examples
func ExampleAdd() {
	sum := learn.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
