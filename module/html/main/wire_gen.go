// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/rema424/twitter-reproduction-backend/lib/dbutil"
	"github.com/rema424/twitter-reproduction-backend/module/html/handler"
	"github.com/rema424/twitter-reproduction-backend/repository"
	"github.com/rema424/twitter-reproduction-backend/service"
)

// Injectors from wire.go:

func InitializeHandlers(db *dbutil.DB) handler.Factory {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	factory := handler.NewFactory(userHandler)
	return factory
}
