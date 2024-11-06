package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func TsetConnext(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}

func Ping(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}
