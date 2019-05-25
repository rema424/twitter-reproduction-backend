package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rema424/twitter-reproduction-backend/model"
)

// UserRepository ...
type UserRepository interface {
	Get(userID int) (*model.User, error)
	SelectAll() ([]*model.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository ...
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

// Get ...
func (r *userRepository) Get(userID int) (*model.User, error) {
	return getUser(r.db, userID)
}

// SelectAll ...
func (r *userRepository) SelectAll() ([]*model.User, error) {
	return selectAllUsers(r.db)
}

func getUser(db *sqlx.DB, userID int) (*model.User, error) {
	q := `SELECT * FROM users WHERE user_id = ?;`
	var u model.User
	if err := db.Get(&u, q, userID); err != nil {
		return nil, err
	}
	return &u, nil
}

func selectAllUsers(db *sqlx.DB) ([]*model.User, error) {
	q := `SELECT * FROM users;`
	var us []*model.User
	if err := db.Select(&us, q); err != nil {
		return nil, err
	}
	return us, nil
}
