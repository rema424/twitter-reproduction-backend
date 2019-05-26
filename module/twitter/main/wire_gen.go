// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/rema424/twitter-reproduction-backend/lib/db"
	"github.com/rema424/twitter-reproduction-backend/module/twitter/handler"
	"github.com/rema424/twitter-reproduction-backend/repository"
	"github.com/rema424/twitter-reproduction-backend/service"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeHandlers(db2 *db.DB) handler.Factory {
	userRepository := repository.NewUserRepository(db2)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	factory := handler.NewFactory(userHandler)
	return factory
}
