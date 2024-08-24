package middleware

import (
	"Glyphrz-go/response"
	"Glyphrz-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAdminAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从上下文获取Claims（JWT负载）
		claims, _ := context.Get("claims")
		// 断言claims是否为Claims结构体指针
		user := claims.(*utils.Claims)
		// Role = 2 代表为管理员
		if user.Role != 1 {
			response.Fail(context, http.StatusForbidden, 403, "该用户没有权限", "")
			context.Abort()
			return
		}
		context.Next()
	}
}
