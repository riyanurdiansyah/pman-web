package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/controller"
	"kalbenutritionals.com/pman/app/helper/model"
)

func InitRoutes(cnf *model.Config) {

	//define gin
	router := gin.Default()

	fmt.Println("CEK : " + cnf.UserSession.SessionKey)
	store := cookie.NewStore([]byte(cnf.UserSession.SessionKey))
	router.Use(sessions.Sessions(cnf.UserSession.SessionID, store))

	authCtrl := &controller.AuthController{}

	cacheConnection := business_logic.NewRedisCacheBL()
	sessionMiddleware := cacheConnection.CheckSession("/")

	router.GET("/signin", sessionMiddleware, authCtrl.Signin)
	router.POST("/signin", authCtrl.Signin)

	AutoGenerateRoutes(router, &controller.MainController{}, "/")

	router.NoRoute(func(ctx *gin.Context) {
		http.StripPrefix("/", http.FileServer(http.Dir("./wwwroot"))).ServeHTTP(ctx.Writer, ctx.Request)
	})

	if err := router.Run(":" + cnf.AppConfig.Port); err != nil {
		panic(err)
	}
}

func AutoGenerateRoutes(router *gin.Engine, controller controller.BaseController, pathPrefix string) {
	cacheConnection := business_logic.NewRedisCacheBL()
	sessionMiddleware := cacheConnection.CheckSession("/signin")

	router.GET(pathPrefix, sessionMiddleware, controller.Index)
	router.GET(pathPrefix+"/:id", sessionMiddleware, controller.Update)
	router.GET(pathPrefix+"/create", sessionMiddleware, controller.Create)
}
