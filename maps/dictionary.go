package maps

import "errors"

type Dictionary map[string]string


var (
	ErrWordExists = errors.New("cannot add word as exists already")
	ErrNotFound =  errors.New("could not find word you are looking for")
)
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
