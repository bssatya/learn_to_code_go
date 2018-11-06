package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just testing"}

	t.Run("Known word", func(t *testing.T) {
		got := dictionary.Search("test")
		want := "this is just testing"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		got := dictionary.Search("search")
		want := ""

		assertStrings(t, got, want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got (%s) Want (%s), Given (%s)", got, want, "test")
	}
}
