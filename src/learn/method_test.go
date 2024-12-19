package learn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dog struct { //nolint:recvcheck
	Name string
	Age  int
}

func (d Dog) Speak() string {
	return "Woof"
}

func (d Dog) FakeBirthday() int {
	// Note: this won't actually increase Dog.Age (not a pointer reference)
	d.Age++
	return d.Age
}

func (d *Dog) Birthday() {
	// *Dog is a "pointer receiver"
	d.Age++
}

func TestDog(t *testing.T) {
	t.Parallel()
	dog := Dog{Name: "Fido", Age: 1}
	assert.Equal(t, "Woof", dog.Speak())
}

func TestDogFakeBirthday(t *testing.T) {
	t.Parallel()
	dog := Dog{Name: "Fido", Age: 1}
	assert.Equal(t, 2, dog.FakeBirthday())
}

func TestDogBirthday(t *testing.T) {
	t.Parallel()
	dog := Dog{Name: "Luna", Age: 1}
	dog.Birthday()
	assert.Equal(t, 2, dog.Age)
}

func ExternalBirthday(d *Dog) {
	d.Age++
}

func TestExternalBirthday(t *testing.T) {
	t.Parallel()
	dog := Dog{Name: "Luna", Age: 1}
	ExternalBirthday(&dog) // pass a dog pointer
	assert.Equal(t, 2, dog.Age)
}
