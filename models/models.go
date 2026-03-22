package models

import "time"

type User struct {
	ID         int64     `db:"id"`
	City       string    `db:"city"`
	Created_at time.Time `db:"created_at"`
}
