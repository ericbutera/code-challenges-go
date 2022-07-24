// https://go.dev/blog/slices-intro
// https://pkg.go.dev/builtin#append

package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	kvs := map[string]string{"a": "apple", "b": "banana"}
	var keys []string
	var vals []string
	for k, v := range kvs {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	assert.ElementsMatch(t, []string{"a", "b"}, keys)
	assert.ElementsMatch(t, []string{"apple", "banana"}, vals)
}
