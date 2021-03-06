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

func TestUpdate(t *testing.T) {
	t.Run("Update the existing word", func(t *testing.T) {
		word := "noun"
		definition := "name of a place"
		dictionary := Dictionary{word: definition}
		newDefinistion := "name of a place or a person"

		err := dictionary.Update(word, newDefinistion)

		assertNoError(t, err)
		asserDefinition(t, dictionary, word, newDefinistion)
	})

	t.Run("Add a new word", func(t *testing.T) {
		word := "verb"
		definition := "represents action"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrorWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {

	word := "noun"
	definition := "name of a place or person"
	dictionary := Dictionary{word: definition}

	t.Run("Delete existing word", func(t *testing.T) {
		err := dictionary.Delete(word)

		_, err = dictionary.Search(word)
		if err != ErrorNotFound {
			t.Errorf("Expected (%s) to be deleted", word)
		}
	})

	t.Run("Delete non existing word", func(t *testing.T) {
		err := dictionary.Delete("not exits")

		if err != ErrorNotFound {
			t.Errorf("Expected (%s) to be deleted", "not exits")
		}
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
