package model

// User ...
type User struct {
	ID          int    `db:"user_id"`
	UserName    string `db:"user_name"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Icon        string `db:"icon"`
	HeaderImage string `db:"header_image"`
	Profile     string `db:"profile"`
	Birthday    string `db:"birthday"`
	Place       string `db:"place"`
	URL         string `db:"url"`
}
