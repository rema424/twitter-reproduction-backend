package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	mydb "github.com/rema424/twitter-reproduction-backend/lib/db"
)

// InitDB ...
func InitDB() *mydb.DB {
	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &mydb.DB{db}
}
