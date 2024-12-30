package db

import "database/sql"

// Queries wraps the generated SQL queries (from sqlc).
type Store interface{
	Querier
}

// NewStore creates and returns a new Store with the provided database connection.
type SQLStore struct {
	db *sql.DB
	*Queries
}
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}