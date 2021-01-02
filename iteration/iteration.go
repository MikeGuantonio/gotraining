package iteration

func Repeat(character string, iterations int) string {
	var result string
	for i := 0; i < iterations; i++ {
		result += character
	}
	return result
}