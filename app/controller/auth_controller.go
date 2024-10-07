package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/helper/constanta"
	"kalbenutritionals.com/pman/app/helper/exception"
)

type AuthController struct {
	AuthBL     *business_logic.AuthBL
	RedisCache *business_logic.RedisCacheBL
}

func NewAuthController(authBL *business_logic.AuthBL, redisCache *business_logic.RedisCacheBL) *AuthController {
	return &AuthController{AuthBL: authBL, RedisCache: redisCache}
}

func (c *AuthController) ChooseRole(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		// cacheConnection := business_logic.NewRedisCacheBL()
		// user := cacheConnection.GetUserLogin(ctx)

		exception.RenderPage(ctx, constanta.AUTH_VIEW_PATH+"choose_role.html", nil, "")
	} else if ctx.Request.Method == "POST" {
	}

}

func (c *AuthController) Signin(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		exception.RenderPage(ctx, constanta.AUTH_VIEW_PATH+"signin.html", nil, "")
	} else if ctx.Request.Method == "POST" {
		session := sessions.Default(ctx)

		session.Options(sessions.Options{
			MaxAge: 3600,
		})

		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		///GET TOKEN JWT FIRST
		token, errToken := c.AuthBL.GetTokenAccess()
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"signin.html", errToken, "Failed to get token")

		///SAVE TOKEN JWT
		session.Set("bearer_token", token)

		data := map[string]string{
			"username": username,
			"password": password,
		}

		body, err := json.Marshal(data)
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"signin.html", err, "Failed to process request")

		headers := map[string]string{
			"Authorization": "Bearer " + session.Get("bearer_token").(string),
		}

		responseLogin, errLogin := c.AuthBL.Login(body, headers)
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"signin.html", errLogin, "Failed connect to server")

		if !responseLogin.BitSuccess {
			exception.RenderPage(ctx, constanta.AUTH_VIEW_PATH+"signin.html", nil, responseLogin.TxtMessage)
		}

		redisName := fmt.Sprintf("%s_%s", responseLogin.ObjData.TxtUserName, responseLogin.TxtGUID)
		session.Set("isLoggedIn", true)
		session.Set("username", username)
		session.Set("session_redis_name", redisName)

		c.RedisCache.Set(redisName, responseLogin)

		errs := session.Save()
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"signin.html", errs, "Failed to save session")

		ctx.Redirect(http.StatusFound, "/")
	}
}
