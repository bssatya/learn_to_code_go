package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("couldnt find the word you are looking for")

func (d Dictionary) Search(key string) (definition string, err error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}
