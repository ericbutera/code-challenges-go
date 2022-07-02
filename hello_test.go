// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Eric")
		want := "Hello, Eric"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
