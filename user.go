package belajargolanggorm

import "time"

type User struct {
	ID        string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
