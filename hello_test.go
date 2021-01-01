package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say Hello To a person", func(t *testing.T) {
		got := Hello("Nick", "")
		want := "Hello Nick"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Default to saying hello if an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Should greet in spanish if language is provided", func(t *testing.T) {
		got := Hello("Nick", "Spanish")
		want := "Hola Nick"
		assertCorrectMessage(t, got, want)
	})
}
