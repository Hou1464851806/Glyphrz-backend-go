package forms

type LoginForm struct {
	Password  string `form:"password" json:"password" binding:"required,min=5,max=20"`
	Name      string `form:"name" json:"name" binding:"required"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaID string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Password string `form:"password" json:"password" binding:"required,min=5,max=20"`
	Name     string `form:"name" json:"name" binding:"required"`
}
