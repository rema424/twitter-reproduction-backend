package dbutil

import (
	"context"
	"log"
	"os"
	"time"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// db ...
type db interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// DB ...
type DB struct {
	*sqlx.DB
}

// NewDB ...
func NewDB() *DB {
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

	return &DB{conn}
}

// Session ...
type Session struct {
	// dbutil.DB インタフェースを満たした dbutil.db をコンポジットする。
	db
	context     context.Context
	warningTime time.Duration
	warningRows int
	secretArgs  bool
}

// NewSession ...
func (db *DB) NewSession(ctx context.Context) *Session {
	if db == nil {
		panic(nil)
	}

	return &Session{
		db:          db.DB,
		context:     ctx,
		warningTime: time.Duration(150) * time.Millisecond,
		warningRows: 500,
	}
}

func (s *Session) check() {
	if s == nil || s.db == nil {
		panic(nil)
	}
}

// Get ...
func (s *Session) Get(dest interface{}, query string, args ...interface{}) error {
	s.check()
	start := time.Now()
	err := s.db.Get(dest, query, args...)
	var rows int
	if err == nil {
		rows = 1
	}
	s.log(query, args, rows, start, err)
	return err
}

// Select ...
func (s *Session) Select(dest interface{}, query string, args ...interface{}) error {
	s.check()
	return s.db.Select(dest, query, args...)
}
