package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (definition string, err error) {
	definition, ok := d[key]
	if !ok {
		return "", errors.New("couldnt find the word you are looking for")
	}
	return definition, nil
}
