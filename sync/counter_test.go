package main

import "testing"

func assertCount(t *testing.T, got Counter, want int){
	t.Helper()
	if (got.Value() != want){
		t.Errorf("Got %d, Want %d", got, want)
	}
}

func TestCounter(t *testing.T){
	t.Run("Should do a basic increment", func(t *testing.T){
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})
}