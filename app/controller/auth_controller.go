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
	AuthBL *business_logic.AuthBL
}

func NewAuthController(authBL *business_logic.AuthBL) *AuthController {
	return &AuthController{AuthBL: authBL}
}

func (c *AuthController) Signin(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		exception.RenderPage(ctx, constanta.AUTH_VIEW_PATH+"signin.html", "")
	} else if ctx.Request.Method == "POST" {
		session := sessions.Default(ctx)

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
			exception.RenderPage(ctx, constanta.AUTH_VIEW_PATH+"signin.html", responseLogin.TxtMessage)
		}

		cacheConnection := business_logic.NewRedisCacheBL()
		redisName := fmt.Sprintf("%s_%s", responseLogin.ObjData.TxtUserName, responseLogin.TxtGUID)
		session.Set("isLoggedIn", true)
		session.Set("username", username)
		session.Set("session_redis_name", redisName)

		cacheConnection.Set(redisName, responseLogin)

		errs := session.Save()
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"signin.html", errs, "Failed to save session")

		ctx.Redirect(http.StatusFound, "/")
	}
}
