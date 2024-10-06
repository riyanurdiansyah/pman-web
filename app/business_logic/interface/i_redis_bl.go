package business_logic

import (
	"github.com/gin-gonic/gin"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
)

type IRedisCachceBL interface {
	Get(key string) ([]byte, error)
	Set(key string, value interface{}) error
	CheckSession(redirect string) gin.HandlerFunc
	GetUserLogin(ctx *gin.Context) *model_response.SigninResponse
}
