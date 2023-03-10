package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

type LoginRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	LoginID       string `json:"login_id,omitempty" valid:"login_id"`
	Password      string `json:"password,omitempty" valid:"password"`
}

// LoginByPhone 通过手机号登录
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required: 手机号必填",
			"digits: 手机号必须是 11 位数字",
		},
		"verify_code": []string{
			"required: 验证码必填",
			"digits: 验证码必须是 6 位数字",
		},
	}
	errs := validate(data, rules, messages)

	login := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(login.Phone, login.VerifyCode, errs)
	return errs
}

// Login 通过手机号登录
func Login(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"login_id": []string{
			"required:登录 ID 为必填项，支持手机号、邮箱和用户名",
			"min:登录 ID 长度需大于 3",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rules, messages)

	login := data.(*LoginRequest)

	errs = validators.ValidateCaptcha(login.CaptchaID, login.CaptchaAnswer, errs)

	return errs
}
