package pg

import (
	"database/sql"

	// this package is imported and blancked because we import it to use the init function.
	_ "github.com/lib/pq"
)

// Store is the struct that represent our PG instance of the store and services.
type Store struct {
	*sql.DB
}

// NewStore returns a new store using postgresql.
// The returned struct could be mainly used to query the database and store / retrieve data on it.
func NewStore(url string) (*Store, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &Store{
		DB: db,
	}, nil
}
