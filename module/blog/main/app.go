package main

import (
	"github.com/rema424/go-gae-blog-app-example/database"
	// "github.com/rema424/go-gae-blog-app-example/module/blog/handler"
	// "github.com/rema424/go-gae-blog-app-example/repository"
	// "github.com/rema424/go-gae-blog-app-example/service"
)

var e = createMux()

func init() {
	// ----
	// DI
	// ----

	// コンストラクタ版
	// db := database.NewDatabase()
	// userRepository := repository.NewUserRepository(db)
	// articleRepository := repository.NewArticleRepository(db)
	// articleService := service.NewArticleService(userRepository, articleRepository)
	// articleHandler := handler.NewArticleHandler(articleService)

	// DIコンテナ版
	// db := database.NewDatabase()
	// articleHandler := InitializeArticleHandler(db)

	// DIコンテナ版 (ファクトリ)
	db := database.NewDatabase()
	handlers := InitializeHandlers(db)

	// --------
	// Routes
	// --------

	// e.GET("/", handler.HelloHandler)
	// e.GET("/:message", handler.ParrotHandler)
	e.GET("/articles/:id", handlers.Article().Show)
}
