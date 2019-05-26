package db

import "github.com/jmoiron/sqlx"

// DB ...
type DB struct {
	*sqlx.DB
}
