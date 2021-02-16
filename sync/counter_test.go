package main

import (
	"testing"
	"sync"
)

func assertCount(t *testing.T, got *Counter, want int){
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

		assertCount(t, &counter, 3)
	})

	t.Run("should run safely concurrently", func(t *testing.T){
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++{
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCount(t, &counter, wantedCount)
	})
}