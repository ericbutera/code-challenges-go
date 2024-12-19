package learn_test

import (
	"testing"

	"github.com/ericbutera/code-challenges-go/src/learn"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	dict := learn.Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		t.Parallel()
		got, _ := dict.Search("test")
		expected := "this is just a test"

		assert.Equal(t, expected, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		t.Parallel()
		_, err := dict.Search("unknown")
		assert.ErrorIs(t, err, learn.ErrNotFound)
	})
}

func TestBuiltIn(t *testing.T) {
	t.Parallel()
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

func TestMake(t *testing.T) {
	t.Parallel()
	bikes := make(map[string]string)
	bikes["yeti"] = "sb130"
	bikes["bmc"] = "teamelite 02 one"
	bikes["cinelli"] = "trackshark"
	assert.Equal(t, "sb130", bikes["yeti"])
	assert.Len(t, 3, len(bikes))
}
