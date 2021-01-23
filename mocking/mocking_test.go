package main

import (
	"bytes"
	"testing"
	"fmt"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spy := &SleeperSpy{}

	Countdown(buffer, spy)

	want := `3
2
1
Go
`
	got := buffer.String()

	if want != got {
		t.Errorf("Wanted %s, Got %s", want, got)
	}

	if spy.Calls != 4 {
		t.Errorf("Wanted %d calls, Got %d calls", 4, spy.Calls)
	}
}
