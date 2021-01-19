package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	definition, found := d[term]

	if !found {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}