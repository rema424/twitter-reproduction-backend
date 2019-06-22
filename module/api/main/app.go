package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/rema424/twitter-reproduction-backend/lib/dbutil"
	"github.com/rema424/twitter-reproduction-backend/module/api/handler"
)

var e = createMux()

func init() {
	e.Use(middleware.CORS())

	// DB の接続は main パッケージで行うことが多いが、
	// AppEngine で複数モジュールが定義されるケースに備えて lib に逃す。
	db := dbutil.NewDB()
	handlers := InitializeHandlers(db)

	e.GET("/", handler.HelloHandler)

	e.GET("/users", handlers.User().Index)
	e.GET("/users/:id", handlers.User().Show)
}
