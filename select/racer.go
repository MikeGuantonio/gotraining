package main

import (
	"time"
	"net/http"
)

func Racer(player1 string, player2 string) (winner string){
	player1Duration := measureResponseTime(player1)
	player2Duration := measureResponseTime(player2)

	if player1Duration < player2Duration {
		return player1
	}
	return player2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}