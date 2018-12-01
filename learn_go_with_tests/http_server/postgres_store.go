package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresPlayerStore struct {
	store *sql.DB
}

func NewPostgresPlayerStore() *PostgresPlayerStore {
	conString := "dbname=playerstore sslmode=disable"
	db, err := sql.Open("postgres", conString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresPlayerStore{store: db}
}

func (p *PostgresPlayerStore) GetLeague() League {
	return nil
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	rows, err := p.store.Query("SELECT wins FROM player_wins WHERE name = $1", name)
	if err != nil {
		return 0
	}
	defer rows.Close()

	for rows.Next() {
		var wins int
		rows.Scan(&wins)
		return wins
	}
	return 0
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	_, err := p.store.Query("INSERT INTO player_wins(name, wins) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET wins = player_wins.wins + 1", name, 1)

	if err != nil {
		fmt.Errorf("Error adding player win to database")
	}
}
