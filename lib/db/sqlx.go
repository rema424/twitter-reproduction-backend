package db

import (
	"log"
	"os"
	"time"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB ...
type DB interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type db struct {
	sqlx *sqlx.DB
}

// NewDB ...
func NewDB() DB {
	conn, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	conn.SetMaxOpenConns(30)
	conn.SetMaxIdleConns(30)
	conn.SetConnMaxLifetime(60 * time.Second)

	return &db{conn}
}

// func (db *DB) Select()

func (db *db) Get(dest interface{}, query string, args ...interface{}) error {
	return db.sqlx.Get(dest, query, args...)
}

func (db *db) Select(dest interface{}, query string, args ...interface{}) error {
	return db.sqlx.Select(dest, query, args...)
}
