//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/rema424/twitter-reproduction-backend/lib/db"
	"github.com/rema424/twitter-reproduction-backend/module/api/handler"
	"github.com/rema424/twitter-reproduction-backend/repository"
	"github.com/rema424/twitter-reproduction-backend/service"
)

func InitializeHandlers(db db.DB) handler.Factory {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
		handler.NewFactory,
	)
	return nil
}
