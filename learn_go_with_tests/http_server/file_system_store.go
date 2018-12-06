package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initialisePlayerDBfile(file)

	if err != nil {
		return nil, fmt.Errorf("Problem initialising player db file")
	}
	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("Problem loading the player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store, err
}

func initialisePlayerDBfile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {

	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}
