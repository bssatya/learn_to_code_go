package main

import "errors"

type Dictionary map[string]string

var (
	ErrorNotFound   = errors.New("couldnt find the word you are looking for")
	ErrorWordExists = errors.New("can not add the word because it already exists")
)

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}
