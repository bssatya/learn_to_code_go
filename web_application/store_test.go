package main

import (
	"database/sql"
	"testing" // the "testify/suite" is used to make the test suite

	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	/*
		The suite is defined as a struct, with the store and db as its attribute. Any variables
		that are to be shared between tests in a suite should be stored as attributes of the
		Suite instance
	*/
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	/*
		The database connection is opened in the setup and stored as an instance variable,
		as in the higher level 'store', that wraps the 'db'
	*/
	connString := "dbname=bird_encyclopedia sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}

	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	/*
		We delete all the entries from the tables before each tests runs, to ensure a
		consistant state before the tests run. In more complex applications, this is
		sometimes achieved in the form of migrations.
	*/

	_, err := s.db.Query("DELETE FROM birds")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	// close the connection after all the tests are run
	s.db.Close()
}

// This is the actual "test" as seem by Go, which runs the tests defined below
func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateBird() {
	s.store.CreateBird(&Bird{
		Description: "test description",
		Species:     "test species",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM birds WHERE description='test description' and species='test species'`)
	if err != nil {
		s.T().Fatal(err)
	}

	// get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Fatal(err)
		}
	}

	// Assert that there must be one entry with the properties of the bird that we just entered(since the data base was empty before we did the insert)
	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuite) TestGetBird() {
	// Insert a sample bird into the 'birds' table
	_, err := s.db.Query(`INSERT INTO birds (species, description) VALUES ('bird', 'description')`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the list of birds through the stores `GetBirds` method
	birds, err := s.store.GetBirds()
	if err != nil {
		s.T().Fatal(err)
	}

	// Assert that the details of the birds is same as the one inserted
	expectedBird := Bird{"bird", "description"}
	if *birds[0] != expectedBird {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedBird, birds[0])
	}
}
