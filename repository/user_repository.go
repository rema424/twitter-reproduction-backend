package repository

import (
	"github.com/rema424/go-gae-blog-app-example/database"
	"github.com/rema424/go-gae-blog-app-example/model"
)

type (
	// UserRepository ...
	UserRepository interface {
		Get(ID int) *model.User
	}

	userRepositoryImpl struct {
		DB database.DB
	}
)

// NewUserRepository ...
func NewUserRepository(DB database.DB) UserRepository {
	return &userRepositoryImpl{
		DB: DB,
	}
}

func (r *userRepositoryImpl) Get(ID int) *model.User {
	return r.DB.User()[ID]
}
