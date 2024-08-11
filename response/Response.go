package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, code int, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}
func Fail(c *gin.Context, httpCode int, code int, msg interface{}, data interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}
