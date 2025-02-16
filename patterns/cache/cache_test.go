package cache_test

import (
	"testing"
	"time"

	"github.com/ericbutera/code-challenges-go/patterns/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test cache set
// cache get

func TestCache(t *testing.T) {
	t.Parallel()

	t.Run("memory string", func(t *testing.T) {
		t.Parallel()
		c := cache.NewInMemoryCache()
		c.Set("key", "value")
		value, ok := c.Get("key")
		require.True(t, ok)
		assert.Equal(t, "value", value)
	})

	t.Run("memory int", func(t *testing.T) {
		t.Parallel()
		c := cache.NewInMemoryCache()
		c.Set("key", 42)
		value, ok := c.Get("key")
		require.True(t, ok)
		assert.Equal(t, 42, value)
	})

	t.Run("memory object", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			ID   string
			Name string
		}
		c := cache.NewInMemoryCache()
		expected := Person{ID: "eric-uuid", Name: "Eric"}
		c.Set(expected.ID, expected)
		value, ok := c.Get(expected.ID)
		require.True(t, ok)
		assert.Equal(t, expected, value)
	})
}

func TestExpiringCache(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		key := "greeting"
		expected := "hello!"
		c := cache.NewExpiringCache[string](5 * time.Minute)
		c.Set(key, expected)
		value, ok := c.Get(key)
		require.True(t, ok)
		assert.Equal(t, expected, *value)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		key := "forty-two"
		expected := 42
		c := cache.NewExpiringCache[int](5 * time.Minute)
		c.Set(key, expected)
		value, ok := c.Get(key)
		require.True(t, ok)
		assert.Equal(t, expected, *value)
	})

	t.Run("object", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			ID   string
			Name string
		}
		expected := Person{ID: "eric-uuid", Name: "Eric"}
		c := cache.NewExpiringCache[Person](5 * time.Minute)
		c.Set(expected.ID, expected)
		value, ok := c.Get(expected.ID)
		require.True(t, ok)
		assert.Equal(t, expected, *value)
	})
}
