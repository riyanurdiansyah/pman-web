package controller

import (
	"github.com/gin-gonic/gin"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/helper/constanta"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
	"kalbenutritionals.com/pman/app/helper/render"
)

type ViewData struct {
	User  model_response.UserData           // Data user object
	Menus []model_response.MenuDataResponse // Data menus berupa list
}

type MainController struct {
	RedisCache *business_logic.RedisCacheBL
}

func NewMainController(redisCache *business_logic.RedisCacheBL) *MainController {
	return &MainController{RedisCache: redisCache}
}

func (hc *MainController) Index(ctx *gin.Context) {
	user := hc.RedisCache.GetUserLogin(ctx)
	menus := hc.RedisCache.GetMenus(ctx)

	viewData := ViewData{
		User:  user.ObjData,
		Menus: menus,
	}

	render.RenderView(ctx, constanta.MAIN_VIEW_PATH+"index.html", viewData)
}
func (hc *MainController) Create(ctx *gin.Context) {
}

func (hc *MainController) Update(ctx *gin.Context) {
	render.RenderView(ctx, constanta.MAIN_VIEW_PATH+"notfound.html", nil)
}

func (hc *MainController) Delete(ctx *gin.Context) {
}
