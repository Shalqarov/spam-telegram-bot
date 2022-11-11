package repository

import (
	"database/sql"
	"os"
)

func SqliteMigration(dsn string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	err = setup(db)

	if err != nil {
		return nil, err
	}

	return db, nil

}

func setup(db *sql.DB) error {
	query, err := os.ReadFile("./internal/repository/migration.sql")

	if err != nil {
		return err
	}

	if _, err := db.Exec(string(query)); err != nil {
		return err
	}

	return nil
}