package model

import "database/sql"

// User ...
type User struct {
	ID          int            `db:"user_id"`
	UserName    string         `db:"user_name"`
	Name        sql.NullString `db:"name"`
	Email       sql.NullString `db:"email"`
	Icon        sql.NullString `db:"icon"`
	HeaderImage sql.NullString `db:"header_image"`
	Profile     sql.NullString `db:"profile"`
	Birthday    sql.NullString `db:"birthday"`
	Place       sql.NullString `db:"place"`
	URL         sql.NullString `db:"url"`
}
