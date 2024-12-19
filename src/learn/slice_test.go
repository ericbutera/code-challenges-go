// https://go.dev/blog/slices-intro
// https://pkg.go.dev/builtin#append

package learn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Parallel()
	kvs := map[string]string{"a": "apple", "b": "banana"}
	keys := make([]string, 0, len(kvs))
	vals := make([]string, 0, len(kvs))
	for k, v := range kvs {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	assert.ElementsMatch(t, []string{"a", "b"}, keys)
	assert.ElementsMatch(t, []string{"apple", "banana"}, vals)
}

func TestAppend(t *testing.T) {
	t.Parallel()
	var a []int
	a = append(a, 1)
	assert.Equal(t, []int{1}, a)
	a = append(a, 2, 3, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, a)
}

func TestAppendLiteral(t *testing.T) {
	t.Parallel()
	data := []int{1}
	data = append(data, 1)
	assert.Equal(t, []int{1, 1}, data)
}

func TestSliceOperator(t *testing.T) {
	t.Parallel()
	bikes := []string{"yeti", "bmc", "cinelli"}
	assert.Equal(t, []string{"yeti"}, bikes[:1]) // omit first = 0
	assert.Equal(t, []string{"yeti", "bmc"}, bikes[0:2])
	assert.Equal(t, []string{"bmc"}, bikes[1:2])
	assert.Equal(t, []string{"bmc", "cinelli"}, bikes[1:3])
	assert.Equal(t, []string{"cinelli"}, bikes[2:]) // omit last = len(bikes)
}

func TestStringSlice(t *testing.T) {
	t.Parallel()
	s := "z̵̼̩̩̿a̴͍̤͍̓́l̷̬̯̓͐ǧ̶͓̫̊̓ͅo̵̧̒͛"
	r := []rune(s)
	assert.Equal(t, []rune{'z', '̵', '̿', '̼', '̩', '̩', 'a', '̴', '̓', '́', '͍', '̤', '͍', 'l', '̷', '̓', '͐', '̬', '̯', 'g', '̶', '̌', '̊', '̓', '͓', 'ͅ', '̫', 'o', '̵', '̒', '͛', '̧'}, r) //nolint:lll
}
