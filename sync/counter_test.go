package main

import "testing"

func TestCounter(t *testing.T){
	counter := Counter{}

	counter.Inc()
	counter.Inc()
	counter.Inc()

	want := 3
	got := counter.Value()
	if (got != want){
		t.Errorf("Got %d, Want %d", got, want)
	}
}