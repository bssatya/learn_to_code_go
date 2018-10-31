package main

// The sql go library is needed to interact with postgres database
import (
	"database/sql"
)

// Our store will have 2 methods, to add a new bird,
// and to get all the existing birds
// each method returns an error, in case if something goes wrong
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

// The 'dbStore' struct will implement the 'Store' interface
// It also takes thr sql DB connection object, which represents
// the database connection
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBird(bird *Bird) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	// THe first underscore means that we don't care about what's returned from
	// this insert query. We just want to know if it was inserted correctly,
	// and the error will be populated if it wasn't
	_, err := store.db.Query("INSERT INTO birds(species, description) VALUES($1, $2)", bird.Species, bird.Description)
	return err
}

func (store *dbStore) GetBirds() ([]*Bird, error) {
	// Query the database for all the birds, and return the result to the row's object
	rows, err := store.db.Query("SELECT species, description FROM birds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the structure that is returned from the function
	// By default, this will be an empty array of birds
	birds := []*Bird{}

	for rows.Next() {
		// for each row returned by the table, create a pointer to bird
		bird := &Bird{}
		// Populate the 'species' and 'description' attributes of the bird,
		// and return incase of the error
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}

		// Finally append the result to the returned array, and repeat for the next row
		birds = append(birds, bird)
	}

	return birds, nil
}

// The store variable is a package level variable that will be available for
// use throught our application code
var store Store

/*
We will need to call the InitStore method to initialize the store. This will typically be done at the begining of our application (in this case, when the server starts up)
This case also be used to set up the store as a mock, which we will be observing later on
*/

func InitStore(s Store) {
	store = s
}
