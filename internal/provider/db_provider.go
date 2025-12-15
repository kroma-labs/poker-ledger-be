package provider

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rotisserie/eris"
)

func provideDB(dbString string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbString)
	if err != nil {
		return nil, eris.Wrap(err, "error opening sqlite connection")
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, eris.Wrap(err, "error pinging sqlite database after opening connection")
	}

	return db, nil
}
