package main 

import (
	"testing"
	"bytes"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	want := "3"
	got := buffer.String()

	if (want != got) {
		t.Errorf("Wanted %s, Got %s", want, got)
	}
}