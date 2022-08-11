package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Speak() string {
	return "Woof"
}

func (d Dog) FakeBirthday() int {
	// Note: this won't actually increase Dog.Age (not a pointer reference)
	d.Age = d.Age + 1
	return d.Age
}

func (d *Dog) Birthday() {
	// *Dog is a "pointer receiver"
	d.Age += 1
}

func TestDog(t *testing.T) {
	dog := Dog{Name: "Fido"}
	assert.Equal(t, "Woof", dog.Speak())
}

func TestDogFakeBirthday(t *testing.T) {
	dog := Dog{Name: "Fido", Age: 1}
	assert.Equal(t, 2, dog.FakeBirthday())
}

func TestDogBirthday(t *testing.T) {
	dog := Dog{Name: "Luna", Age: 1}
	dog.Birthday()
	assert.Equal(t, 2, dog.Age)
}

func ExternalBirthday(d *Dog) {
	d.Age += 1
}

func TestExternalBirthday(t *testing.T) {
	dog := Dog{Name: "Luna", Age: 1}
	ExternalBirthday(&dog) // pass a dog pointer
	assert.Equal(t, 2, dog.Age)
}
