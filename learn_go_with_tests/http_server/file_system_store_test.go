package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from the reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertNoError(t, err)
		assertLeague(t, got, want)
	})

	t.Run("Get Player store", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")

		want := 33

		assertNoError(t, err)
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for the existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertNoError(t, err)
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for the new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertNoError(t, err)
		assertScoreEquals(t, got, want)
	})

	t.Run("Works with empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
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
