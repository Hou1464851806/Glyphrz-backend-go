package initialize

import (
	"Glyphrz-go/middleware"
	"Glyphrz-go/router"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 路由组
	ApiGroup := r.Group("/v1/")
	router.UserRouter(ApiGroup)
	r.Use(middleware.Logger(), middleware.Recovery(true))
	return r
}
