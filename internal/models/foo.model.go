package models

type Foo struct {
	Title       string `db:"title" json:"title"`
	Director    string `db:"director" json:"director"`
	ReleaseYear string `db:"release_year" json:"release_year"`
}
