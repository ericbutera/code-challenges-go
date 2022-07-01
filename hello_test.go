// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
