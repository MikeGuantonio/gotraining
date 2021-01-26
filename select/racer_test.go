package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Should find the faster website", func(t *testing.T) {
		slowServer := createServer(20 * time.Millisecond)
		fastServer := createServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		
		want := fastURL
		got, _:= Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("Wanted %v, Got %v", want, got)
		}
	})

	t.Run("Should return an error if requests take longer than 10 seconds", func(t *testing.T) {
		slowServer := createServer(11 * time.Second)
		fastServer := createServer(12 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		
		_, err := Racer(slowURL, fastURL)

		if err == nil {
			t.Errorf("Wanted a timeout error but nothing was returned")
		}
	})
	
}

func createServer(delay time.Duration) *httptest.Server{
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}