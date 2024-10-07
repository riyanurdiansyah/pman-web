package middleware

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"kalbenutritionals.com/pman/app/controller"
	"kalbenutritionals.com/pman/app/helper/model"
	"kalbenutritionals.com/pman/app/injector"
)

func InitRoutes(cnf *model.Config) {

	//define gin
	router := gin.Default()

	redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	store := cookie.NewStore([]byte(cnf.UserSession.SessionKey))
	router.Use(sessions.Sessions(cnf.UserSession.SessionID, store))

	authCtrl, err := injector.InitializeAuthController()
	if err != nil {
		log.Fatalf("Failed to initialize auth controller: %v", err)
	}

	redisCache, err := injector.InitializeRedisCacheBL()
	if err != nil {
		log.Fatalf("Failed to initialize auth controller: %v", err)
	}
	sessionMiddleware := redisCache.CheckSession("/")

	router.GET("/signin", sessionMiddleware, authCtrl.Signin)
	router.POST("/signin", sessionMiddleware, authCtrl.Signin)

	router.GET("/choose-role", authCtrl.ChooseRole)
	router.POST("/choose-role", authCtrl.ChooseRole)

	AutoGenerateRoutes(router, &controller.MainController{}, "/")

	router.NoRoute(func(ctx *gin.Context) {
		http.StripPrefix("/", http.FileServer(http.Dir("./wwwroot"))).ServeHTTP(ctx.Writer, ctx.Request)
	})

	if err := router.Run(":" + cnf.AppConfig.Port); err != nil {
		panic(err)
	}
}

func AutoGenerateRoutes(router *gin.Engine, controller controller.BaseController, pathPrefix string) {
	redisCache, err := injector.InitializeRedisCacheBL()
	if err != nil {
		log.Fatalf("Failed to initialize auth controller: %v", err)
	}
	sessionMiddleware := redisCache.CheckSession("/signin")

	router.GET(pathPrefix, sessionMiddleware, controller.Index)
	router.GET(pathPrefix+"/:id", sessionMiddleware, controller.Update)
	router.GET(pathPrefix+"/create", sessionMiddleware, controller.Create)
}
