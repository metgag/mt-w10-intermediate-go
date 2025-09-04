package models

import "time"

type Movie struct {
	ID          uint16    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Director    string    `db:"director" json:"director"`
	ReleaseYear string    `db:"release_year" json:"release_year"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
