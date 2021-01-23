package main

import (
	"os"
	"io"
	"fmt"
)

func Countdown(writer io.Writer) {
	for _, number := range([]int{3, 2, 1}) {
		fmt.Fprintf(writer, "%d\n", number)	
	}
	fmt.Fprintf(writer, "Go")
}

func main() {
	Countdown(os.Stdout)
}