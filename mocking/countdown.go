package main

import (
	"os"
	"io"
	"fmt"
)

func Countdown(writer io.Writer) {
	fmt.Fprintf(writer, "3")
}

func main() {
	Countdown(os.Stdout)
}