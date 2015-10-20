package models

import (
	"time"
)

type Member struct {
	// General information.
	Id					int64		`db:"id"`
	Name				string		`db:"name"`
	Email				string		`db:"email"`
	Password			[]byte		`db:"password"`
	RegistrationDate	time.Time	`db:"registration_date"`

	// SRS-specific.
	Level 				int32		`db:"level"`
}