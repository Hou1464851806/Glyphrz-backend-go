package utils

import (
	"Glyphrz-go/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateToken(c *gin.Context, ID uint, name string, role uint) string {
	// jwt和payload信息
	j := NewJWT()
	claims := Claims{
		ID:   ID,
		Name: name,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60, // 1 hour
		},
	}
	// 生成token
	tokenString, err := j.CreateToken(claims)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, 500, "token生成错误", "")
		return ""
	}
	return tokenString
}
