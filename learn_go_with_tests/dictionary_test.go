package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just testing"}

	got := Search(dictionary, "test")
	want := "this is just testing"

	if got != want {
		t.Errorf("Got (%s) Want (%s), Given (%s)", got, want, "test")
	}
}
