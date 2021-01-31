package main

import (
	"net/http"
	"time"
	"fmt"
)

var TEN_SECOND_TIMEOUT = 10 * time.Second

func Racer(player1 string, player2 string) (winner string, err error){
	winner, err = ConfigurableRacer(player1, player2, TEN_SECOND_TIMEOUT)
	return winner, err
}

func ConfigurableRacer(player1 string, player2 string, timeout time.Duration) (winner string, err error){
	select {
	case <- ping(player1):
			return player1, nil
	case <- ping(player2):
			return player2, nil
	case <- time.After(timeout):
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