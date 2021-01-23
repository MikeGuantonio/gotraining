package main

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nick")

	got := buffer.String()
	want := "Hello Nick"

	if (want != got) {
		t.Errorf("Wanted %s, Got %s", want, got)
	}
}