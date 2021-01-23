package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	for _, number := range []int{3, 2, 1} {
		fmt.Fprintf(writer, "%d\n", number)
	}
	fmt.Fprintf(writer, "Go")
}

func main() {
	Countdown(os.Stdout)
}
