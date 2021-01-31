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
		got, err:= Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Was not expecting error got %v", err)
		}

		if got != want {
			t.Errorf("Wanted %v, Got %v", want, got)
		}
	})
}

func TestConfigurableRacer(t *testing.T) {
	t.Run("Should return an error if requests take longer than 10 seconds", func(t *testing.T) {
		server := createServer(25 * time.Millisecond)

		defer server.Close()
		
		_, err := ConfigurableRacer(server.URL, server.URL, 20 * time.Millisecond)

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