package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassicForLoop(t *testing.T) {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += 1
	}
	assert.Equal(t, sum, 3)
}

func TestDecrement(t *testing.T) {
	sum := 3
	i := 3
	for i > 0 {
		sum -= 1
		i--
	}
	assert.Equal(t, sum, 0)
}

func TestRangeOverArray(t *testing.T) {
	sum := 0
	for _, v := range []int{1, 2, 3} {
		sum += v
	}
	assert.Equal(t, sum, 6)
}

func TestRangeOverMap(t *testing.T) {
	// note: maps aren't ordered!
	data := map[string]string{
		"salsa": "mukluk",
		"bmc":   "roadmachine 105",
	}
	outputs := map[string]string{}

	for k, v := range data {
		outputs[k] = k + ":" + v
	}

	assert.Equal(t, outputs["salsa"], "salsa:mukluk")
	assert.Equal(t, outputs["bmc"], "bmc:roadmachine 105")
}
