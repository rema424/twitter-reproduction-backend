//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/rema424/twitter-reproduction-backend/database"
	"github.com/rema424/twitter-reproduction-backend/module/twitter/handler"
	"github.com/rema424/twitter-reproduction-backend/repository"
	"github.com/rema424/twitter-reproduction-backend/service"
)

func InitializeHandlers(db database.DB) handler.Factory {
	wire.Build(
		repository.NewUserRepository,
		repository.NewArticleRepository,
		service.NewArticleService,
		handler.NewArticleHandler,
		handler.NewFactory,
	)
	return nil
}
