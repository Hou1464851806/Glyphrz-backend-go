package controller

import (
	"Glyphrz-go/DAO"
	"Glyphrz-go/forms"
	"Glyphrz-go/global"
	"Glyphrz-go/model"
	"Glyphrz-go/response"
	"Glyphrz-go/utils"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login
// 用户登录
func Login(c *gin.Context) {
	// 字段校验
	loginForm := forms.LoginForm{}
	if err := c.ShouldBind(&loginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	//// 图形验证码检验
	//if !store.Verify(loginForm.CaptchaID, loginForm.Captcha, true) {
	//	response.Fail(c, http.StatusBadRequest, 400, "验证码错误", "")
	//	return
	//}
	// 查询用户是否注册过
	user, ok := DAO.FindUserByNameDao(loginForm.Name)
	if !ok {
		response.Fail(c, http.StatusUnauthorized, 401, "用户未注册", "")
		return
	}
	// 检验用户信息
	if user.Password != loginForm.Password {
		response.Fail(c, http.StatusUnauthorized, 401, "用户名或密码错误", "")
		return
	}
	// 登录成功，生成token
	token := utils.CreateToken(c, user.ID, user.Name, user.Role)
	// 通过设置cookie存储token并返回user给客户端
	c.SetCookie("token", token, 60*60*24*365*10, "/", "", false, true)
	response.Success(c, 200, "success", user)
}

// GetUserList
// 获取用户列表
func GetUserList(c *gin.Context) {
	userList := DAO.GetUserListDao()
	response.Success(c, 200, "success", map[string]interface{}{
		"users": userList,
	})
}

func Register(c *gin.Context) {
	// 字段校验
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 查询用户是否注册过
	_, ok := DAO.FindUserByNameDao(registerForm.Name)
	if ok {
		response.Fail(c, http.StatusBadRequest, 400, "用户已存在", "")
		return
	}
	// 新建用户
	user := &model.User{
		Name:     registerForm.Name,
		Password: registerForm.Password,
	}
	ok = DAO.AddUser(user)
	if !ok {
		response.Fail(c, http.StatusInternalServerError, 500, "新建用户错误", "")
	}
	response.Success(c, http.StatusOK, "success", "")
}

func Logout(c *gin.Context) {
	// 获取负载
	claims, _ := c.Get("claims")
	user := claims.(*utils.Claims)
	// 将jti存储在Redis中，并且过期时间和JWT一致
	err := global.Redis.Set(user.Id, "1", time.Unix(user.ExpiresAt, 0).Sub(time.Now())).Err()
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, 500, "黑名单添加失败", "")
		return
	}
	color.Blue(fmt.Sprint(user.StandardClaims))
	response.Success(c, 200, "用户已登出", "")
	return
}
