//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	"github.com/rema424/twitter-reproduction-backend/module/twitter/handler"
	"github.com/rema424/twitter-reproduction-backend/repository"
	"github.com/rema424/twitter-reproduction-backend/service"
)

func InitializeHandlers(db *sqlx.DB) handler.Factory {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
		handler.NewFactory,
	)
	return nil
}
