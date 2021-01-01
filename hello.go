package hello

import (
	"fmt"
)

const prefixInEnglish = "Hello"

func Hello(name string) string {
	if "" == name { 
		return prefixInEnglish
	}
	return prefixInEnglish + " " + name
}

func main() {
	fmt.Println(Hello("Nick"))
}