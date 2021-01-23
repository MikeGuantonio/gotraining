package main

import (
	"os"
	"io"
	"fmt"
)

func Countdown(writer io.Writer) {
	fmt.Fprintf(writer, "3 ")
	fmt.Fprintf(writer, "2 ")
	fmt.Fprintf(writer, "1 ")
	fmt.Fprintf(writer, "Go")
}

func main() {
	Countdown(os.Stdout)
}