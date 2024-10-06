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
	authDAL := data_access.NewAuthDAL()
	authBL := business_logic.NewAuthBL(authDAL)
	authController := controller.NewAuthController(authBL)
	return authController, nil
}

// injector.go:

var SuperSet = wire.NewSet(data_access.NewAuthDAL, business_logic.NewAuthBL, controller.NewAuthController)
