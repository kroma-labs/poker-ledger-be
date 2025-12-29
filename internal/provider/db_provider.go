package provider

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // DB driver
	"github.com/rotisserie/eris"
)

func provideDB(dbString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", dbString)
	if err != nil {
		return nil, eris.Wrap(err, "error connecting to sqlite DB")
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
