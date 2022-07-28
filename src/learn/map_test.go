package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		expected := "this is just a test"

		assert.Equal(t, expected, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		assert.Error(t, err, ErrNotFound)
	})
}

func TestBuiltIn(t *testing.T) {
	kvs := map[string]string{"a": "apple", "b": "banana"}
	keys := ""
	vals := ""
	for k, v := range kvs {
		keys += k
		vals += v
	}
	assert.Equal(t, "ab", keys)
	assert.Equal(t, "applebanana", vals)
}
