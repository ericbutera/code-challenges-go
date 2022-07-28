// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Eric")
		assert.Equal(t, "Hello, Eric", got)
	})
	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		assert.Equal(t, "Hello, World", got)
	})
}
