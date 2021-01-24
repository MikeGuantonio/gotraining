package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalCountdown = "Go"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	countdownNumbers := []int{3, 2, 1}

	for _, number := range countdownNumbers {
		sleeper.Sleep()
		fmt.Fprintf(writer, "%d\n", number)
	}

	sleeper.Sleep()
	fmt.Fprintf(writer, "%s\n", finalCountdown)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
