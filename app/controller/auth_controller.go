package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	business_logic "kalbenutritionals.com/pman/app/business_logic"
	business_logic_i "kalbenutritionals.com/pman/app/business_logic/interface"
	"kalbenutritionals.com/pman/app/helper/constanta"
	"kalbenutritionals.com/pman/app/helper/exception"
)

type AuthController struct {
	AuthBL     business_logic_i.IAuthBL
	RedisCache *business_logic.RedisCacheBL
}

func NewAuthController(authBL business_logic_i.IAuthBL, redisCache *business_logic.RedisCacheBL) *AuthController {
	return &AuthController{AuthBL: authBL, RedisCache: redisCache}
}

func (c *AuthController) ChooseRole(ctx *gin.Context) {
	user := c.RedisCache.GetUserLogin(ctx)

	if ctx.Request.Method == "GET" {
		exception.RenderPage(ctx, constanta.AUTH_VIEW_PATH+"choose_role.html", user.ObjData, "")
	} else if ctx.Request.Method == "POST" {
		session := sessions.Default(ctx)

		selectedRole := ctx.PostForm("role")

		fmt.Println("CEK ROLE : " + selectedRole)

		data := map[string]string{
			"intRoleID":   selectedRole,
			"txtUserName": user.ObjData.TxtUserName,
		}

		headers := map[string]string{
			"Authorization": "Bearer " + session.Get("bearer_token").(string),
		}

		body, err := json.Marshal(data)
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"signin.html", err, "Failed to process request")

		menus, errAPI := c.AuthBL.GetMenus(body, headers)
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"choose-role.html", errAPI, "Failed connect to server")

		redisNameMenu := fmt.Sprintf("%s_%s_menu", user.ObjData.TxtUserName, user.TxtGUID)

		c.RedisCache.Set(redisNameMenu, menus)

		ctx.Redirect(http.StatusFound, "/")
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

		token, errToken := c.AuthBL.GetTokenAccess()
		exception.HandleError(ctx, constanta.AUTH_VIEW_PATH+"signin.html", errToken, "Failed to get token")

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

		if len(responseLogin.ObjData.LtRoles) > 1 {
			ctx.Redirect(http.StatusFound, "/choose-role")
		} else {
			ctx.Redirect(http.StatusFound, "/")
		}
	}
}
