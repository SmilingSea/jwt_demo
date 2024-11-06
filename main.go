package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"jwt_demo/mw"
	"jwt_demo/routes"
)

func main() {

	h := server.Default()
	mw.InitJwt()

	// 加载路由
	routes.LoadRoutes(h)
	h.Spin()
}
