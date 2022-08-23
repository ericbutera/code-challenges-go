package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	var value int = 10
	var pointy *int = &value
	assert.Equal(t, value, 10)
	assert.Equal(t, *pointy, 10)
	assert.Equal(t, pointy, &value)
	assert.Equal(t, *pointy, value)
}
