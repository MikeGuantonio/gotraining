package hello

import (
	"fmt"
)

const prefixInEnglish = "Hello"
const prefixInSpanish = "Hola"
const prefixInFrench = "Bonjour"

const spanish = "Spanish"
const french = "French"

// const languages = map[string]string {
// 	"Spanish": "Hola", 
// 	"French": "Bonjour",
// 	"English": "Hello",
// }

func Hello(name string, language string) string {
	if "" == name {
		return prefixInEnglish
	}

	return greetingPrefix(language) + " " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = prefixInSpanish
	case french: 
		prefix = prefixInFrench
	default:
		prefix = prefixInEnglish
	}
	return prefix
}

// func greetingPrefix2(language string) (prefix string) {
// 	prefix, found := language[language]
// 	if !found {
// 		prefix = language["English"]
// 	}
// 	return prefix
// }

func main() {
	fmt.Println(Hello("Nick", "English"))
}
