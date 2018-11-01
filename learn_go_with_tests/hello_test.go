package main

import "testing"

func Test_hello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s', Want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := hello("Satya", "Kannada")
		want := "Hello, Satya"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := hello("Satya", "Spanish")
		want := "Hola, Satya"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := hello("Satya", "French")
		want := "Bonjour, Satya"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := hello("", "French")
		want := "Bonjour, World"

		assertCorrectMessage(t, got, want)
	})
}
