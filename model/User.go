package model

// User ...
type User struct {
	ID    int    `db:"user_id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
