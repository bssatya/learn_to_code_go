package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("couldnt find the word you are looking for")

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
