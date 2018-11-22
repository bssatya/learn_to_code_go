package main

import (
	"database/sql"

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

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	return 12
}

func (p *PostgresPlayerStore) RecordWin(name string) {

}
