package main

import (
	"net/http"
)

func Racer(player1 string, player2 string) (winner string){
	select {
	case <- ping(player1):
			return player1
	case <- ping(player2):
			return player2
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