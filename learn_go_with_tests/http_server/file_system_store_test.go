package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from the reader", func (t *testing.T){
		database := strings.NewReader(`[{"name": "Cleo", "Wins": 10}, {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player {
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})
}
