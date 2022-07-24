// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration
package learn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	assert.Equal(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	nap := Repeat("z", 3)
	fmt.Println(nap)
	// Output: zzz
}
