package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	value := d[term]

	if value == "" {
		return "", errors.New("could not find the word you were looking for")
	}

	return value, nil
}