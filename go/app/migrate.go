package app

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

// migrate runs migration from specified folder
func migrate(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error applying postgres driver for migrations: %w", err)
	}

	return goose.Up(db, "./migrations")
}
