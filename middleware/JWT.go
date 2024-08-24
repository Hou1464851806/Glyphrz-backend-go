package middleware

import (
	"Glyphrz-go/global"
	"Glyphrz-go/response"
	"Glyphrz-go/utils"
	"errors"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从请求的Cookie的token字段获得用户token
		tokenString, err := context.Cookie("token")
		if err != nil {
			color.Red("token cookie not found")
			response.Fail(context, http.StatusUnauthorized, 401, "请先登录", "")
			context.Abort()
			return
		}
		if tokenString == "" {
			response.Fail(context, http.StatusUnauthorized, 401, "请先登录", "")
			context.Abort()
			return
		}
		color.Yellow(tokenString)
		// 解析token
		j := utils.NewJWT()
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.Fail(context, http.StatusUnauthorized, 401, "token已过期", "")
				context.Abort()
				return
			}
			response.Fail(context, http.StatusUnauthorized, 401, "token验证错误", "")
			context.Abort()
			return
		}
		// 查询JWT是否在Redis黑名单中
		color.Blue(claims.Id)
		if jti, _ := global.Redis.Get(claims.Id).Result(); jti != "" {
			color.Red(claims.Id)
			response.Fail(context, http.StatusUnauthorized, 401, "token已失效", "")
			context.Abort()
			return
		}
		// 记录claims和userID，给下游使用
		//color.Blue(fmt.Sprint(context))
		context.Set("claims", claims)
		context.Set("userID", claims.ID)
		context.Next()
	}
}
