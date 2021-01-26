package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL
	
	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("Wanted %v, Got %v", want, got)
	}
}