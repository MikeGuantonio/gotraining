package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	want := `3
2
1
Go`
	got := buffer.String()

	if want != got {
		t.Errorf("Wanted %s, Got %s", want, got)
	}
}
