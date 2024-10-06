//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/controller"
	"kalbenutritionals.com/pman/app/data_access"
)

var SuperSet = wire.NewSet(
	data_access.NewAuthDAL,
	business_logic.NewAuthBL,
	controller.NewAuthController,
)

func InitializeAuthController() (*controller.AuthController, error) {
	wire.Build(SuperSet)
	return &controller.AuthController{}, nil // Placeholder, akan diisi oleh Wire
}
