package main

import (
	"github.com/rema424/twitter-reproduction-backend/lib/db"
	"github.com/rema424/twitter-reproduction-backend/lib/echoutil"
	"github.com/rema424/twitter-reproduction-backend/module/html/handler"
)

var e = createMux()

func init() {
	echoutil.SetCommonMiddleware(e)

	// templates := tmplutil.NewTemplateContainer()
	// templates["hello"] = tmplutil.RegisterTemplate("layout", "hello")
	// e.Renderer = tmplutil.NewEchoRenderer(templates)

	// DB の接続は main パッケージで行うことが多いが、
	// AppEngine で複数モジュールが定義されるケースに備えて lib に逃す。
	db := db.NewDB()
	handlers := InitializeHandlers(db)

	e.GET("/", handler.HelloHandler)

	e.GET("/users", handlers.User().Index)
	e.GET("/users/:id", handlers.User().Show)
}
