package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := createServer(20 * time.Millisecond)
	fastServer := createServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL
	
	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("Wanted %v, Got %v", want, got)
	}
}

func createServer(delay time.Duration) *httptest.Server{
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}