package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just testing"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just testing"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("search")
		want := ErrorNotFound

		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "noun"
		definition := "name of a place or a person"
		err := dictionary.Add(word, definition)
		assertNoError(t, err)
		asserDefinition(t, dictionary, word, definition)
	})

	t.Run("Add a existing word", func(t *testing.T) {
		word := "noun"
		definition := "name of a place or a person"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "name of anything")
		assertError(t, err, ErrorWordExists)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got (%s) Want (%s), Given (%s)", got, want, "test")
	}
}

func asserDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != definition {
		t.Errorf("Got '%s', Want '%s'", got, definition)
	}
}
