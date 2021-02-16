package main

import (
	"testing"
	"http/testing"
)

type StubStore struct {
	response string
}

func(*ss StubStore) Fetch() string {
	return ss.response
}

func TestServer(t *testing.T) {
	data := "hello halavanja"
	store := StubStore{response: data}
	server := Server(&store)

	request := httpTest.NewRequest(http.MethodGet, "/", nil)
	response := httpTest.NewRecorder()

	server.ServeHttp(response, request)

	if response.Body.String() != data {
		t.Errorf("Wanted %s, Got %s", data, response.Body.String())
	}
}
