package routes

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"jwt_demo/handlers"
	"jwt_demo/mw"
)

func LoadRoutes(h *server.Hertz) {
	h.POST("/login", mw.JwtMiddleware.LoginHandler)
	// 测试连通
	h.GET("/ping", handlers.TsetConnext)
	auth := h.Group("/auth", mw.JwtMiddleware.MiddlewareFunc(), mw.RateLimitMiddleware(1, 1))
	auth.GET("/ping", handlers.Ping)

}
