package models

type Login struct {
	Email    string `db:"email" json:"email" validate:"email"`
	Password string `db:"password" json:"password"`
}
