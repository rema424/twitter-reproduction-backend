//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/rema424/go-gae-blog-app-example/database"
	"github.com/rema424/go-gae-blog-app-example/module/blog/handler"
	"github.com/rema424/go-gae-blog-app-example/repository"
	"github.com/rema424/go-gae-blog-app-example/service"
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
