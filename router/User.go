package router

import "github.com/gin-gonic/gin"

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "user-list",
			})
		})
		UserRouter.GET("panic", func(c *gin.Context) {
			panic("test recovery middleware")
		})
	}
}
