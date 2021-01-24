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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
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
	cSleeper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep: time.Sleep,
	}
	Countdown(os.Stdout, cSleeper)
}
