package learn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	t.Parallel()
	value := 10
	pointy := &value
	assert.Equal(t, 10, value)
	assert.Equal(t, 10, *pointy)
	assert.Equal(t, pointy, &value)
	assert.Equal(t, *pointy, value)
}
