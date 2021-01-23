package main

import (
	"fmt"
	"io"
	"os"
)

const finalCountdown = "Go"

func Countdown(writer io.Writer) {
	countdownNumbers := []int{3, 2, 1}

	for _, number := range countdownNumbers {
		fmt.Fprintf(writer, "%d\n", number)
	}
	fmt.Fprintf(writer, finalCountdown)
}

func main() {
	Countdown(os.Stdout)
}
