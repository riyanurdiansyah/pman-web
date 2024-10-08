package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kalbenutritionals.com/pman/app/helper/constanta"
	"kalbenutritionals.com/pman/app/middleware"
)

type Server struct {
	Router *gin.Engine
}

func InitEnvironment() {
	env := os.Getenv("GO_ENV")
	var envFile string

	if env == "dev" {
		envFile = ".env.dev"
	} else if env == "prod" {
		envFile = ".env.prod"
	} else {
		fmt.Println("Unknown environment, defaulting to .env.dev")
		envFile = ".env.dev"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
}

func Run() {
	InitEnvironment()

	cnf := constanta.Get()

	middleware.InitRoutes(&cnf)
}
