package learn_test

import (
	"cmp"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://pkg.go.dev/cmp
func TestMultiColumnSort(t *testing.T) {
	t.Parallel()

	type Order struct {
		Product  string
		Customer string
		Price    float64
	}

	expected := []Order{
		{"foo", "alice", 2.00},
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"foo", "bob", 4.00},
		{"bar", "carol", 1.00},
		{"baz", "carol", 4.00},
	}

	orders := []Order{
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"baz", "carol", 4.00},
		{"foo", "alice", 2.00},
		{"bar", "carol", 1.00},
		{"foo", "bob", 4.00},
	}
	// Sort by customer first, product second, and last by higher price
	slices.SortFunc(orders, func(a, b Order) int {
		return cmp.Or(
			strings.Compare(a.Customer, b.Customer),
			strings.Compare(a.Product, b.Product),
			cmp.Compare(b.Price, a.Price),
		)
	})
	assert.Equal(t, expected, orders)
}
