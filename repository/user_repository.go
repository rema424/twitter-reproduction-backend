package repository

import (
	"context"

	"github.com/rema424/twitter-reproduction-backend/lib/dbutil"
	"github.com/rema424/twitter-reproduction-backend/model"
)

// UserRepository ...
type UserRepository interface {
	Get(ctx context.Context, userID int) (*model.User, error)
	SelectAll(ctx context.Context) ([]*model.User, error)
}

type userRepository struct {
	db *dbutil.DB
}

// NewUserRepository ...
func NewUserRepository(db *dbutil.DB) UserRepository {
	return &userRepository{db}
}

// Get ...
func (r *userRepository) Get(ctx context.Context, userID int) (*model.User, error) {
	return getUser(ctx, r.db, userID)
}

func getUser(ctx context.Context, db *dbutil.DB, userID int) (*model.User, error) {
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
	if err := db.NewSession(ctx).Get(&u, q, userID); err != nil {
		return nil, err
	}
	return &u, nil
}

// SelectAll ...
func (r *userRepository) SelectAll(ctx context.Context) ([]*model.User, error) {
	return selectAllUsers(ctx, r.db)
}

func selectAllUsers(ctx context.Context, db *dbutil.DB) ([]*model.User, error) {
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
	if err := db.NewSession(ctx).Select(&us, q); err != nil {
		return nil, err
	}
	return us, nil
}
