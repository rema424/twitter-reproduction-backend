package main

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
	db := InitDB()
	handlers := InitializeHandlers(db)

	// --------
	// Routes
	// --------

	// e.GET("/", handler.HelloHandler)
	// e.GET("/:message", handler.ParrotHandler)
	e.GET("/users", handlers.User().Index)
	e.GET("/users/:id", handlers.User().Show)
}
