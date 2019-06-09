package main

import (
	"github.com/rema424/twitter-reproduction-backend/lib/db"
	echoutil "github.com/rema424/twitter-reproduction-backend/lib/echo"
	templateutil "github.com/rema424/twitter-reproduction-backend/lib/template"
	"github.com/rema424/twitter-reproduction-backend/module/html/handler"
)

var e = createMux()

func init() {
	echoutil.SetCommonMiddleware(e)

	templates := templateutil.NewTemplateContainer()
	templates["hello"] = templateutil.RegisterTemplate("layout", "hello")
	e.Renderer = templateutil.NewEchoRenderer(templates)

	// DB の接続は main パッケージで行うことが多いが、
	// AppEngine で複数モジュールが定義されるケースに備えて lib に逃す。
	db := db.NewDB()
	handlers := InitializeHandlers(db)

	e.GET("/", handler.HelloHandler)

	e.GET("/users", handlers.User().Index)
	e.GET("/users/:id", handlers.User().Show)
}
