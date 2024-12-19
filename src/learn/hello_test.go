// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package learn_test

import (
	"testing"

	"github.com/ericbutera/code-challenges-go/src/learn"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Parallel()
	t.Run("saying hello to people", func(t *testing.T) {
		t.Parallel()
		got := learn.Hello("Eric")
		assert.Equal(t, "Hello, Eric", got)
	})
	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		t.Parallel()
		got := learn.Hello("")
		assert.Equal(t, "Hello, World", got)
	})
}

func TestTableDrivenHello(t *testing.T) {
	t.Parallel()
	// https://github.com/golang/go/wiki/TableDrivenTests#example-of-a-table-driven-test
	tests := []struct {
		in  string
		out string
	}{
		{"Eric", "Hello, Eric"},
		{"", "Hello, World"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			got := learn.Hello(tt.in)
			assert.Equal(t, tt.out, got)
		})
	}
}

func TestTableDrivenMapHello(t *testing.T) {
	t.Parallel()
	// https://github.com/golang/go/wiki/TableDrivenTests#using-a-map-to-store-test-cases
	tests := map[string]struct {
		input  string
		result string
	}{
		"greets name":                    {"Eric", "Hello, Eric"},
		"greets world unless name given": {"", "Hello, World"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := learn.Hello(test.input)
			assert.Equal(t, test.result, got)
		})
	}
}
