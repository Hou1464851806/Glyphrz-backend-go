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
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	r.Use(middleware.Logger(), middleware.Recovery(true))
	return r
}
