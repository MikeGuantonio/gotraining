package hello

import (
	"testing"
)

func TestShouldReturnHelloPhrase(t *testing.T) {
	got := Hello("Nick")
	want := "Hello Nick"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}