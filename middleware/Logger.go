package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 路径
		path := c.Request.URL.Path
		// 请求
		query := c.Request.URL.RawQuery
		// 调用下一个handler
		c.Next()
		// 计算耗时
		cost := time.Since(start)
		// 记录异常日志
		if c.Writer.Status() != 200 {
			zap.L().Info(
				path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}
	}
}
