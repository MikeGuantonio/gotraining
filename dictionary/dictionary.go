package dictionary

type Dictionary map[string]string

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already present. not updated")
)

func (d Dictionary) Search(term string) (string, error) {
	definition, found := d[term]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(term string, definition string) error {
	_, err := d.Search(term)

	// I don't like this switch in the tutorial
	// for some cases it returns and others it
	// will do a thing.
	// It should only return a value and let the function
	// do its work 01-22-2021
	switch err {
	case ErrNotFound:
		d[term] = definition
		break
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
