package hello

import (
	"fmt"
)

const prefixInEnglish = "Hello"
const prefixInSpanish = "Hola"
const prefixInFrench = "Bonjour"

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if "" == name {
		return prefixInEnglish
	}

	prefix := prefixInEnglish

	switch language {
	case spanish:
		prefix = prefixInSpanish
	case french: 
		prefix = prefixInFrench
	default:
		prefix = prefixInEnglish
	}

	return prefix + " " + name
}

func main() {
	fmt.Println(Hello("Nick", "English"))
}
