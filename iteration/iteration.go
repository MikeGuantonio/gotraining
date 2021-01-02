package iteration

func Repeat(character string) string {
	var result string
	for i := 0; i < 5; i++ {
		result = result + character
	}
	return result
}