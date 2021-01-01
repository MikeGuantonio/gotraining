package hello

import (
	"fmt"
)

const prefixInEnglish = "Hello"
const prefixInSpanish = "Hola"

const spanish = "Spanish"

func Hello(name string, language string) string {
	if "" == name {
		return prefixInEnglish
	}

	prefix := prefixInEnglish

	switch language {
	case spanish:
		prefix = prefixInSpanish
	default:
		prefix = prefixInEnglish
	}

	return prefix + " " + name
}

func main() {
	fmt.Println(Hello("Nick", "English"))
}
