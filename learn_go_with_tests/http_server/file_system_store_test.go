package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from the reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})

	t.Run("Get Player store", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")

		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for the existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for the new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
