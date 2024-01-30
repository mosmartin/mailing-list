package db

import (
	"time"
)

type Email struct {
	ID          int64
	Email       string
	ConfirmedAt *time.Time
	OptOut      bool
}
