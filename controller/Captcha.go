package controller

import (
	"Glyphrz-go/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, 500, "验证码生成错误", "")
	}
	response.Success(c, 200, "success", map[string]interface{}{
		"captchaID": id,
		"image":     b64s,
	})
}
