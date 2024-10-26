// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/google/wire"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/controller"
	"kalbenutritionals.com/pman/app/data_access"
)

// Injectors from injector.go:

func InitializeAuthController() (*controller.AuthController, error) {
	iAuthDAL := data_access.NewAuthDAL()
	iAuthBL := business_logic.NewAuthBL(iAuthDAL)
	client := business_logic.ProvideRedisClient()
	redisCacheBL := business_logic.ProvideRedisCacheBL(client)
	authController := controller.NewAuthController(iAuthBL, redisCacheBL)
	return authController, nil
}

func InitializeMainController() (*controller.MainController, error) {
	client := business_logic.ProvideRedisClient()
	redisCacheBL := business_logic.ProvideRedisCacheBL(client)
	mainController := controller.NewMainController(redisCacheBL)
	return mainController, nil
}

func InitializeRedisCacheBL() (*business_logic.RedisCacheBL, error) {
	client := business_logic.ProvideRedisClient()
	redisCacheBL := business_logic.ProvideRedisCacheBL(client)
	return redisCacheBL, nil
}

// injector.go:

var authSet = wire.NewSet(data_access.NewAuthDAL, business_logic.NewAuthBL, business_logic.ProvideRedisClient, business_logic.ProvideRedisCacheBL, controller.NewAuthController)

var mainSet = wire.NewSet(business_logic.ProvideRedisClient, business_logic.ProvideRedisCacheBL, controller.NewMainController)

var redisSet = wire.NewSet(business_logic.ProvideRedisClient, business_logic.ProvideRedisCacheBL)
