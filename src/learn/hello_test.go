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

func TestTableDrivenHello(t *testing.T) {
	// https://github.com/golang/go/wiki/TableDrivenTests#example-of-a-table-driven-test
	var hellotests = []struct {
		in  string
		out string
	}{
		{"Eric", "Hello, Eric"},
		{"", "Hello, World"},
	}
	for _, tt := range hellotests {
		t.Run(tt.in, func(t *testing.T) {
			got := Hello(tt.in)
			assert.Equal(t, tt.out, got)
		})
	}
}

func TestTableDrivenMapHello(t *testing.T) {
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
			got := Hello(test.input)
			assert.Equal(t, test.result, got)
		})
	}
}
