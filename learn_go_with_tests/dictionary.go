package main

type Dictionary map[string]string

func (d Dictionary) Search(dictionary map[string]string, key string) string {
	return d[key]
}
