package main

import "github.com/rema424/twitter-reproduction-backend/module/twitter/handler"

var e = createMux()

func init() {
	db := InitDB()
	handlers := InitializeHandlers(db)

	e.GET("/", handler.HelloHandler)

	e.GET("/users", handlers.User().Index)
	e.GET("/users/:id", handlers.User().Show)
}
