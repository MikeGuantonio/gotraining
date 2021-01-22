package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(term string) (string, error) {
	definition, found := d[term]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(term string, definition string) {
	d[term] = definition
}