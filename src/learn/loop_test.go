package learn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassicForLoop(t *testing.T) {
	t.Parallel()
	sum := 0
	for i := 0; i < 3; i++ { //nolint:intrange
		sum++
	}
	assert.Equal(t, 3, sum)
}

func TestDecrement(t *testing.T) {
	t.Parallel()
	sum := 3
	i := 3
	for i > 0 {
		sum--
		i--
	}
	assert.Equal(t, 0, sum)
}

func TestRangeOverArray(t *testing.T) {
	t.Parallel()
	sum := 0
	for _, v := range []int{1, 2, 3} {
		sum += v
	}
	assert.Equal(t, 6, sum)
}

func TestRangeOverMap(t *testing.T) {
	t.Parallel()
	// note: maps aren't ordered!
	data := map[string]string{
		"salsa": "mukluk",
		"bmc":   "roadmachine 105",
	}
	outputs := map[string]string{}

	for k, v := range data {
		outputs[k] = k + ":" + v
	}

	assert.Equal(t, "salsa:mukluk", outputs["salsa"])
	assert.Equal(t, "bmc:roadmachine 105", outputs["bmc"])
}
