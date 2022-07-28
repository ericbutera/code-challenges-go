// https://gobyexample.com/arrays
package learn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	var a [5]int
	a[4] = 100
	assert.Equal(t, 0, a[0])
	assert.Equal(t, 100, a[4])

	b := [5]int{1, 2, 3, 4, 5}
	assert.Equal(t, 1, b[0])
	assert.Equal(t, 5, b[len(b)-1])
}

func Test2dArray(t *testing.T) {
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	res := fmt.Sprint(twoD)
	assert.Equal(t, "[[0 1 2] [1 2 3]]", res)
}
