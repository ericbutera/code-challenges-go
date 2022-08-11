package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func TestPerson(t *testing.T) {
	person := Person{Name: "Luna", Age: 42}
	assert.Equal(t, "Luna", person.Name)
	assert.Equal(t, 42, person.Age)
}

func TestIdiomaticNewPerson(t *testing.T) {
	luna := newPerson("Luna", 4)
	assert.Equal(t, "Luna", luna.Name)
	assert.Equal(t, 4, luna.Age)
}
