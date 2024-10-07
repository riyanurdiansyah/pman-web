//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/controller"
	"kalbenutritionals.com/pman/app/data_access"
)

var authSet = wire.NewSet(
	data_access.NewAuthDAL,
	business_logic.NewAuthBL,
	business_logic.ProvideRedisClient,
	business_logic.ProvideRedisCacheBL,
	controller.NewAuthController,
)

var mainSet = wire.NewSet(
	business_logic.ProvideRedisClient,
	business_logic.ProvideRedisCacheBL,
	controller.NewMainController,
)

var redisSet = wire.NewSet(
	business_logic.ProvideRedisClient,
	business_logic.ProvideRedisCacheBL,
)

func InitializeAuthController() (*controller.AuthController, error) {
	wire.Build(authSet)
	return &controller.AuthController{}, nil // Placeholder, akan diisi oleh Wire
}

func InitializeMainController() (*controller.MainController, error) {
	wire.Build(mainSet)
	return &controller.MainController{}, nil
}

func InitializeRedisCacheBL() (*business_logic.RedisCacheBL, error) {
	wire.Build(redisSet)
	return &business_logic.RedisCacheBL{}, nil
}
