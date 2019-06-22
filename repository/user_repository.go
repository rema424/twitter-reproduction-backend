package repository

import (
	"github.com/rema424/twitter-reproduction-backend/lib/dbutil"
	"github.com/rema424/twitter-reproduction-backend/model"
)

// UserRepository ...
type UserRepository interface {
	Get(userID int) (*model.User, error)
	SelectAll() ([]*model.User, error)
}

type userRepository struct {
	db dbutil.DB
}

// NewUserRepository ...
func NewUserRepository(db dbutil.DB) UserRepository {
	return &userRepository{db}
}

// Get ...
func (r *userRepository) Get(userID int) (*model.User, error) {
	return getUser(r.db, userID)
}

func getUser(db dbutil.DB, userID int) (*model.User, error) {
	q := `
	SELECT
		user_id,
		created,
		updated,
		user_name,
		COALESCE(name, "") AS name,
		COALESCE(email, "") AS email,
		COALESCE(icon, "") AS icon,
		COALESCE(header_image, "") AS header_image,
		COALESCE(profile, "") AS profile,
		COALESCE(birthday, "") AS birthday,
		COALESCE(place, "") AS place,
		COALESCE(url, "") AS url
	FROM users
	WHERE user_id = ?;
	`
	var u model.User
	if err := db.Get(&u, q, userID); err != nil {
		return nil, err
	}
	return &u, nil
}

// SelectAll ...
func (r *userRepository) SelectAll() ([]*model.User, error) {
	return selectAllUsers(r.db)
}

func selectAllUsers(db dbutil.DB) ([]*model.User, error) {
	q := `
	SELECT
		user_id,
		created,
		updated,
		user_name,
		COALESCE(name, "") AS name,
		COALESCE(email, "") AS email,
		COALESCE(icon, "") AS icon,
		COALESCE(header_image, "") AS header_image,
		COALESCE(profile, "") AS profile,
		COALESCE(birthday, "") AS birthday,
		COALESCE(place, "") AS place,
		COALESCE(url, "") AS url
	FROM users;
	`
	var us []*model.User
	if err := db.Select(&us, q); err != nil {
		return nil, err
	}
	return us, nil
}
