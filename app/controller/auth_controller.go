package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"kalbenutritionals.com/pman/app/business_logic"
	"kalbenutritionals.com/pman/app/helper/api"
	"kalbenutritionals.com/pman/app/helper/constanta"
	model_request "kalbenutritionals.com/pman/app/helper/model/request"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
)

type AuthController struct {
}

func (c *AuthController) Signin(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		RenderPage(ctx, "")
	} else if ctx.Request.Method == "POST" {
		SigninPOST(ctx)
	}
}

func RenderPage(ctx *gin.Context, errorMsg string) {
	tmpl, err := template.ParseFiles(constanta.AUTH_VIEW_PATH + "signin.html")
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	err = tmpl.Execute(ctx.Writer, map[string]string{"ErrorMessage": errorMsg})
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
}

func SigninPOST(ctx *gin.Context) {
	var tokenResponse model_request.TokenResponse
	var signinResponse model_response.SigninResponse
	session := sessions.Default(ctx)
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	data := map[string]string{
		"username": username,
		"password": password,
	}

	body, err := json.Marshal(data)
	if err != nil {
		RenderPage(ctx, err.Error())
		return
	}
	// token := session.Get("bearer_token").(string)
	// if token == "" {
	// var err error
	response, err := api.PostRefreshToken(constanta.TOKEN_URL, nil)
	if err != nil {
		RenderPage(ctx, err.Error())
		return
	}
	errJson := json.Unmarshal(response, &tokenResponse)
	if errJson != nil {
		RenderPage(ctx, errJson.Error())
		return
	}
	session.Set("bearer_token", tokenResponse.AccessToken)
	// }
	res, err := api.PostRequest(constanta.LOGIN_URL, body, map[string]string{
		"Authorization": "Bearer " + session.Get("bearer_token").(string),
	})
	if err != nil {
		RenderPage(ctx, err.Error())
		return
	}

	errJsonSignin := json.Unmarshal(res, &signinResponse)
	if errJson != nil {
		RenderPage(ctx, errJsonSignin.Error())
		return
	}
	if !signinResponse.BitSuccess {
		RenderPage(ctx, signinResponse.TxtMessage)
		return
	}

	cacheConnection := business_logic.NewRedisCacheBL()
	redisName := fmt.Sprintf("%s_%s", signinResponse.ObjData.TxtUserName, signinResponse.TxtGUID)
	session.Set("isLoggedIn", true)
	session.Set("username", username)
	session.Set("session_redis_name", redisName)

	cacheConnection.Set(redisName, signinResponse)

	errs := session.Save()
	if errs != nil {
		RenderPage(ctx, errs.Error())
		return
	}
	ctx.Redirect(http.StatusFound, "/")
}
