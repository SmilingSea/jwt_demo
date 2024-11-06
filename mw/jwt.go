package mw

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
	limiterMap    = make(map[string]*rate.Limiter)
	mu            sync.Mutex
)

// 限制访问频率，防止恶意刷接口
func RateLimitMiddleware(rps float64, burst int) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取客户端IP
		clientIP := c.ClientIP()
		mu.Lock()
		limiter, exists := limiterMap[clientIP]
		if !exists {
			// 创建一个新的速率限制器，限制每秒 rps 次请求，允许 burst 次突发请求
			limiter = rate.NewLimiter(rate.Limit(rps), burst)
			limiterMap[clientIP] = limiter
		}
		mu.Unlock()

		// 检查请求速率
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, utils.H{
				"code":    http.StatusTooManyRequests,
				"message": "too many requests, please try again later",
			})
			c.Abort()
			return
		}
	}
}

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "HS256",
		Key:              []byte("demo"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		TokenLookup:      "header:Authorization, query: token, cookie: jwt",
		TokenHeadName:    "Bearer",
		// 登录成功后的响应
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
		},
		// 收到登录数据后的处理逻辑
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Username string `json:"account"`
				Password string `json:"password"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			username := loginStruct.Username
			password := loginStruct.Password
			if !(username == "admin" && password == "admin") {
				return nil, errors.New("invalid username or password")
			}
			return username, nil
		},
		IdentityKey: IdentityKey,
		// 从 token 提取用户信息的函数
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			return nil
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
