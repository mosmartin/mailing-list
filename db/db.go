package db

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/mattn/go-sqlite3"
)

type Email struct {
	ID          int64
	Email       string
	ConfirmedAt *time.Time
	OptOut      bool
}

func CreateDB(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE emails (
			id            INTEGER PRIMARY KEY,
			email         TEXT UNIQUE,
			confirmed_at  INTEGER,
			opt_out       INTEGER
		);
	`)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			// error if table already exists
			// if sqliteErr.Code != sqlite3.ErrConstraint {
			// 	return err
			// }

			if sqliteErr.Code != 1 {
				slog.Error(sqliteErr.Error())
			}
		} else {
			slog.Error("Error creating emails table: %s", err)
			
			return err
		}
	}

	return nil
}

func getEmail(row *sql.Row) (*Email, error) {
	email := &Email{}
	err := row.Scan(&email.ID, &email.Email, &email.ConfirmedAt, &email.OptOut)
	if err != nil {
		slog.Error("Error scanning email: %s", err)

		return nil, err
	}

	// convert time from unix timestamp to time.Time
	t := time.Unix(email.ConfirmedAt.Unix(), 0)
	email.ConfirmedAt = &t

	return email, nil
}