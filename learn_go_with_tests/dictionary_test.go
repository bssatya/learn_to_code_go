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
	dictionary := Dictionary{}

	t.Run("Add a word", func(t *testing.T) {
		dictionary.Add("noun", "name of a place or a person")
		want := "name of a place or a person"
		got, err := dictionary.Search("noun")

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if want != got {
			t.Errorf("Got %s, Want %s", got, want)
		}
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got (%s) Want (%s), Given (%s)", got, want, "test")
	}
}
