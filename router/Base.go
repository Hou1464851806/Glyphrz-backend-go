package router

import (
	"Glyphrz-go/controller"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", controller.GetCaptcha)
	}
}
