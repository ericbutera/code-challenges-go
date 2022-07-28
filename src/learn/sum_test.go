package learn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	assert.Equal(t, 15, got)
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	assert.Equal(t, 4, sum)
}

// https://go.dev/blog/examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
