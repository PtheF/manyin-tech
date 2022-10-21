package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-service/service"
	"go-service/utils"
)

type AccountController struct {
	BaseController
}

// VerifyAccount post {cp, pw, tp = 'cp' || 'pw', cpt = "p" || "e"}
func (c *AccountController) VerifyAccount() {
	tp := c.GetString("tp")
	pw := c.GetString("pw")
	cp := c.GetString("cp")

	if tp == "" ||
		(tp != "pw" && tp != "cp") ||
		(tp == "cp" && cp == "") ||
		(tp == "pw" && pw == "") {
		c.Resp(ParamIllegal)
	}

	user := c.GetUser()

	if err := service.VerifyService(cp, pw, tp, user); err != nil {
		logs.Error("verify error: %v", err)

		c.Resp(VerifyError)
	}

	c.Success("验证成功")

}

func (c *AccountController) UpdateAccount() {
	nike := c.GetString("nikeName")
	phone := c.GetString("phone")
	email := c.GetString("email")

	user := c.GetUser()

	logs.Info("update account params: nike=%v, phone=%v, email=%v", nike, phone, email)

	if nike == "" || len(nike) > 50 ||
		(user.Phone != "" && (phone != "" && user.Phone != phone)) ||
		(user.Email != "" && (email != "" && user.Email != email)) ||
		(phone != "" && utils.RegCheck(phone, utils.PHONE_REG) != 1) ||
		(email != "" && utils.RegCheck(email, utils.MAIL_REG) != 1) ||
		(nike == "" && phone == "" && email == "") {
		c.Resp(ParamIllegal)
	}

	user.NikeName = nike
	if user.Phone == "" {
		user.Phone = phone
	}

	if user.Email == "" {
		user.Email = email
	}

	if err := service.UpdateAccount(user); err != nil {
		c.Failed(err.Error())
	}

	if sign, err := c.GenJwt(user); err != nil {
		c.Failed(err.Error())
	} else {
		c.ReturnData(sign)
	}
}

func (c *AccountController) UpdatePhone() {
	phone := c.GetString("phone")
	cp := c.GetString("cap")
	user := c.GetUser()

	if phone == user.Phone || cp == "" {
		c.Resp(ParamIllegal)
	}

	user.Phone = phone

	if err := service.UpdateAccountWithVerify(phone, cp, user); err != nil {
		c.Failed(err.Error())
	}

	if sign, err := c.GenJwt(user); err != nil {
		c.Failed(err.Error())
	} else {
		c.ReturnData(sign)
	}
}

func (c *AccountController) UpdateEmail() {
	email := c.GetString("email")
	cp := c.GetString("cap")
	user := c.GetUser()

	user.Email = email

	if err := service.UpdateAccountWithVerify(email, cp, user); err != nil {
		c.Failed(err.Error())
	}

	if sign, err := c.GenJwt(user); err != nil {
		c.Failed(err.Error())
	} else {
		c.ReturnData(sign)
	}
}

func (c *AccountController) UpdatePassword() {
	pw := c.GetString("pw")
	user := c.GetUser()

	if err := service.UpdatePassword(pw, user.Uuid); err != nil {
		c.Failed(err.Error())
	}

	c.Success("修改成功")
}
