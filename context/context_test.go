package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (ss *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return ss.response
}

func (ss *SpyStore) Cancel() {
	ss.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("Should send data from server", func(t *testing.T) {
		data := "hello halavanjaa"
		store := SpyStore{response: data}
		server := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("Wanted %s, Got %s", data, response.Body.String())
		}

		if store.cancelled {
			t.Errorf("Request should not have been cancelled due to timeout")
		}
	})

	t.Run("Should cancel work for long running process", func(t *testing.T) {
		data := "hello halavanja"
		store := SpyStore{response: data}
		server := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Create a context that will cancel after a time
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("Request should have been cancelled due to timeout")
		}
	})

}
