package main

import (
	"net/http"
	"time"
	"fmt"
)

func Racer(player1 string, player2 string) (winner string, err error){
	select {
	case <- ping(player1):
			return player1, nil
	case <- ping(player2):
			return player2, nil
	case <- time.After(10 * time.Second):
			return "", fmt.Errorf("Time out. 10 seconds exceeded")
	}
}

func ping(url string) chan struct{} { 
	channel := make(chan struct{})
	go func() {
		http.Get(url)
		close(channel)
	}()
	return channel
}