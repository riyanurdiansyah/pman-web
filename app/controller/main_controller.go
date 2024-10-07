package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/helper/constanta"
	"kalbenutritionals.com/pman/app/helper/exception"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
	"kalbenutritionals.com/pman/app/helper/render"
)

type MainController struct {
	RedisCache *business_logic.RedisCacheBL
}

func NewMainController(redisCache *business_logic.RedisCacheBL) *MainController {
	return &MainController{RedisCache: redisCache}
}

func (hc *MainController) Index(ctx *gin.Context) {
	var signinResponse model_response.SigninResponse
	session := sessions.Default(ctx)
	sessionRedis := session.Get("session_redis_name").(string)
	if sessionRedis == "" {
		ctx.Redirect(http.StatusUnauthorized, constanta.SIGNIN)
	}
	val, err := hc.RedisCache.Get(sessionRedis)
	exception.HandleErrorRedirect(ctx, constanta.SIGNIN, err)

	errs := json.Unmarshal(val, &signinResponse)
	exception.HandleErrorRedirect(ctx, constanta.SIGNIN, errs)

	render.RenderView(ctx, constanta.MAIN_VIEW_PATH+"index.html", nil)
}
func (hc *MainController) Create(ctx *gin.Context) {
}

func (hc *MainController) Update(ctx *gin.Context) {
	render.RenderView(ctx, constanta.MAIN_VIEW_PATH+"notfound.html", nil)
}

func (hc *MainController) Delete(ctx *gin.Context) {
}
