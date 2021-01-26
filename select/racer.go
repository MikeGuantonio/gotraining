package main

import (
	"time"
	"net/http"
)

func Racer(player1 string, player2 string) (winner string){
	startPlayer1 := time.Now()
	http.Get(player1)
	player1Duration := time.Since(startPlayer1)

	startPlayer2 := time.Now()
	http.Get(player2)
	player2Duration := time.Since(startPlayer2)

	if player1Duration < player2Duration {
		return player1
	}
	return player2
}