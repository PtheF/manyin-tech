package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-service/service"
	"go-service/utils"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	id := c.GetString("id")
	pw := c.GetString("pw")

	logs.Info("user login: id=%v, pw=%v", id, pw)
	tp := utils.RegCheck(id, utils.PHONE_REG, utils.MAIL_REG)

	if tp == 0 || utils.RegCheck(pw, utils.PASS_REG) == 0 {
		c.Resp(ParamIllegal)
	}

	userDTO, loginError := service.Login(id, pw, tp)

	logs.Info("login userDTO: %v", userDTO)

	if loginError != nil {
		logs.Info("login error: %v", loginError)
		if loginError.DatabaseError {
			c.Resp(DatabaseError)
		} else if loginError.NoSuchUser || loginError.PasswordIncorrect {
			c.Resp(UserLoginError)
		}
	}

	sign, err := c.GenJwt(userDTO)

	if err != nil {
		c.Failed(err.Error())
	}

	c.SetJWT(sign)

	c.ReturnData(userDTO)
}

func (c *UserController) Register() {
	id := c.GetString("id")
	pw := c.GetString("pw")
	cp := c.GetString("cap")

	logs.Info("register:id=%v, pw=%v, cp=%v", id, pw, cp)

	// 1. 检查数据合法
	i := utils.RegCheck(id, utils.PHONE_REG, utils.MAIL_REG)
	if i == 0 || utils.RegCheck(pw, utils.PASS_REG) == 0 || cp == "" {
		logs.Info("i=%v", i)
		c.Resp(ParamIllegal)
	}

	// 2. 检查验证码
	if _, err := service.NewCaptcha(id).Verify(cp); err != nil {
		c.Failed(err.Error())
	}

	// 3. 注册
	if err := service.Register(id, pw); err != nil {
		c.Resp(UserExist)
	}

	c.Success("注册成功")
}

// Captcha 发送验证码
func (c *UserController) Captcha() {
	id := c.Ctx.Input.Query("id")

	if utils.RegCheck(id, utils.PHONE_REG, utils.MAIL_REG) == 0 {
		c.Resp(ParamIllegal)
	}

	if captcha, err := service.NewCaptcha(id).Send(); err != nil {
		c.Failed(err.Error())
	} else {
		c.ReturnData(captcha)
	}
}
