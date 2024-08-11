package router

import (
	"Glyphrz-go/controller"
	"Glyphrz-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.Login)
		UserRouter.GET("list", middleware.JWTAuth(), controller.GetUserList)
		UserRouter.POST("register", controller.Register)
		UserRouter.GET("panic", func(c *gin.Context) {
			panic("test recovery middleware")
		})
	}
}
