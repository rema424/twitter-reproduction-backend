package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// InitDB ...
func InitDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
