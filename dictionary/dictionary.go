package dictionary

func Search(dictionary map[string]string, term string) string {
	return dictionary[term]
}