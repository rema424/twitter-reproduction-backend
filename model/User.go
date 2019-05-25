package model

import "time"

// User ...
type User struct {
	ID          int       `db:"user_id"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
	UserName    string    `db:"user_name"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Icon        string    `db:"icon"`
	HeaderImage string    `db:"header_image"`
	Profile     string    `db:"profile"`
	Birthday    string    `db:"birthday"`
	Place       string    `db:"place"`
	URL         string    `db:"url"`
}
