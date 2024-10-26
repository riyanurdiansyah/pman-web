package business_logic

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"kalbenutritionals.com/pman/app/helper/constanta"
	"kalbenutritionals.com/pman/app/helper/exception"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
)

// RedisCacheBL struct
type RedisCacheBL struct {
	Client *redis.Client
	Ctx    context.Context
}

// Constructor untuk RedisCacheBL
func NewRedisCacheBL(client *redis.Client) *RedisCacheBL {
	return &RedisCacheBL{
		Client: client,
		Ctx:    context.Background(),
	}
}

// Provider untuk redis.Client
func ProvideRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// Provider untuk RedisCacheBL
func ProvideRedisCacheBL(client *redis.Client) *RedisCacheBL {
	return NewRedisCacheBL(client)
}

// GetUserLoggin implements business_logic.IRedisCachceBL.
func (r *RedisCacheBL) GetUserLogin(ctx *gin.Context) *model_response.SigninResponse {
	var signinResponse model_response.SigninResponse

	session := sessions.Default(ctx)
	sessionRedis := session.Get("session_redis_name").(string)
	if sessionRedis == "" {
		ctx.Redirect(http.StatusUnauthorized, constanta.SIGNIN)
	}
	val, err := r.Get(sessionRedis)
	exception.HandleErrorRedirect(ctx, constanta.SIGNIN, err)

	errs := json.Unmarshal(val, &signinResponse)
	exception.HandleErrorRedirect(ctx, constanta.SIGNIN, errs)

	return &signinResponse
}

func (r *RedisCacheBL) GetMenus(ctx *gin.Context) []model_response.MenuDataResponse {
	var menus []model_response.MenuDataResponse

	session := sessions.Default(ctx)
	sessionRedis := session.Get("session_redis_name").(string)
	if sessionRedis == "" {
		ctx.Redirect(http.StatusUnauthorized, constanta.SIGNIN)
	}
	sessionRedisMenu := sessionRedis + "_menu"
	val, err := r.Get(sessionRedisMenu)
	exception.HandleErrorRedirect(ctx, constanta.SIGNIN, err)

	errs := json.Unmarshal(val, &menus)
	exception.HandleErrorRedirect(ctx, constanta.SIGNIN, errs)

	return menus
}

func RedirectTo(ctx *gin.Context, redirect string) {
	currentPath := ctx.Request.URL.Path
	if strings.Contains(currentPath, "signin") {
		ctx.Next()
	} else {
		ctx.Redirect(http.StatusSeeOther, redirect)
		ctx.Abort()
		return
	}
}

func (r *RedisCacheBL) CheckSession(redirect string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentPath := ctx.Request.URL.Path
		session := sessions.Default(ctx)
		sessionRedis, ok := session.Get("session_redis_name").(string)
		if !ok {
			if strings.Contains(currentPath, "signin") {
				ctx.Next()
				return
			} else {
				ctx.Redirect(http.StatusSeeOther, redirect)
				ctx.Abort()
				return
			}

		} else {
			if sessionRedis == "" {
				RedirectTo(ctx, redirect)
			} else {
				val, err := r.Get(sessionRedis)
				exception.HandleErrorRedirect(ctx, constanta.SIGNIN, err)

				if val == nil {
					RedirectTo(ctx, redirect)
				} else {
					if strings.Contains(currentPath, "signin") {
						ctx.Redirect(http.StatusSeeOther, redirect)
						ctx.Abort()
						return
					}
				}
			}
		}

	}
}

// Get implements business_logic.IRedisCachceBL.
func (r *RedisCacheBL) Get(key string) ([]byte, error) {
	val, err := ProvideRedisClient().Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

// Set implements business_logic.IRedisCachceBL.
func (r *RedisCacheBL) Set(key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(context.Background(), key, jsonData, time.Duration(60)*time.Minute).Err()
}
