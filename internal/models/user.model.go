package models

type User struct {
	ID       uint16 `db:"id" json:"id"`
	Email    string `db:"email" json:"email" validate:"email"`
	Password string `db:"password" json:"password" validate:"min=8,containsany=!@#$%^&*,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}
