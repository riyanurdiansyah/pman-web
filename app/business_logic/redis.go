package business_logic

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	business_logic "kalbenutritionals.com/pman/app/business_logic/interface"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
)

type RedisCacheBL struct {
	rdb *redis.Client
}

// GetUserLoggin implements business_logic.IRedisCachceBL.
func (r *RedisCacheBL) GetUserLogin(ctx *gin.Context) *model_response.SigninResponse {
	var signinResponse model_response.SigninResponse
	session := sessions.Default(ctx)
	sessionRedis := session.Get("session_redis_name").(string)
	val, err := r.Get(sessionRedis)
	if err != nil {
		ctx.Redirect(http.StatusUnauthorized, "/")
	}

	json.Unmarshal(val, &signinResponse)

	return &signinResponse
}

func (r *RedisCacheBL) CheckSession(redirect string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if session.Get("isLoggedIn") == nil {
			ctx.Redirect(http.StatusSeeOther, redirect)
			ctx.Abort()
			return
		}
		ctx.Next()
	}

	// return func(ctx *gin.Context) {
	// 	session := sessions.Default(ctx)
	// 	if session.Get("isLoggedIn") == nil {
	// 		if strings.Contains(redirect, "signin") {
	// 			ctx.Redirect(http.StatusSeeOther, redirect)
	// 			ctx.Abort()
	// 			return
	// 		} else {
	// 			ctx.Next()
	// 		}
	// 	} else {
	// 		if strings.Contains(redirect, "signin") {
	// 			ctx.Next()
	// 		} else {
	// 			ctx.Redirect(http.StatusSeeOther, redirect)
	// 			ctx.Abort()
	// 			return
	// 		}
	// 	}
	// }
}

func NewRedisCacheBL() business_logic.IRedisCachceBL {
	// addr, err := helper.Decrypt(os.Getenv("ReddisConnection"), []byte(os.Getenv("RijndaelKey")))
	// if err != nil {
	// 	log.Fatal("Error decrypting connection string Reddis " + err.Error())
	// }
	return &RedisCacheBL{rdb: redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})}
}

// Get implements business_logic.IRedisCachceBL.
func (r *RedisCacheBL) Get(key string) ([]byte, error) {
	val, err := r.rdb.Get(context.Background(), key).Result()
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
	return r.rdb.Set(context.Background(), key, jsonData, time.Duration(60)*time.Minute).Err()
}
