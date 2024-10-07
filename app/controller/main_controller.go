package controller

import (
	"github.com/gin-gonic/gin"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/helper/constanta"
	"kalbenutritionals.com/pman/app/helper/render"
)

type MainController struct {
	RedisCache *business_logic.RedisCacheBL
}

func NewMainController(redisCache *business_logic.RedisCacheBL) *MainController {
	return &MainController{RedisCache: redisCache}
}

func (hc *MainController) Index(ctx *gin.Context) {
	user := hc.RedisCache.GetUserLogin(ctx)

	render.RenderView(ctx, constanta.MAIN_VIEW_PATH+"index.html", user.ObjData)
}
func (hc *MainController) Create(ctx *gin.Context) {
}

func (hc *MainController) Update(ctx *gin.Context) {
	render.RenderView(ctx, constanta.MAIN_VIEW_PATH+"notfound.html", nil)
}

func (hc *MainController) Delete(ctx *gin.Context) {
}
