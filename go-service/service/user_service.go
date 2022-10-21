package service

import (
	"database/sql"
	"github.com/astaxie/beego/logs"
	uuid "github.com/satori/go.uuid"
	"go-service/models"
	"go-service/utils"
	"strings"
	"time"
)

type UserError struct {
	Msg               string
	NoSuchUser        bool
	PasswordIncorrect bool
	DatabaseError     bool
	IllegalParam      bool
	RepeatId          bool
}

func (e *UserError) Error() string {
	return e.Msg
}

// Login
/*
  tp = 1: phone
  tp = 2: email
*/
func Login(id string, pw string, tp int) (userDTO *models.UserDTO, userError *UserError) {

	var user *models.UserPO

	var err error

	if tp == 1 {
		user, err = models.GetUserByPhone(id)
	} else if tp == 2 {
		user, err = models.GetUserByEmail(id)
	} else {
		return userDTO, &UserError{IllegalParam: true, Msg: "非法参数"}
	}

	// 没有用户
	if err != nil {
		logs.Error("dao error: %v", err)
		return userDTO, &UserError{NoSuchUser: true, Msg: err.Error()}
	}

	encodePw, err := utils.EncodeWithSalt(pw, user.Salt)
	// 密码加密错误
	if err != nil {
		logs.Error("password encode error: %v, user{login=%v, password=%v}", err, id, pw)
		return userDTO, &UserError{PasswordIncorrect: true, Msg: err.Error()}
	}

	// 密码不匹配
	if encodePw != user.Password {
		logs.Info("password incorrect: user{login=%v, password=%v}", id, pw)
		return userDTO, &UserError{PasswordIncorrect: true, Msg: "密码错误"}
	}

	logs.Info("login dao user: %v", user)

	return &models.UserDTO{
		Uuid:     user.Uuid,
		Email:    user.Email,
		Phone:    user.Phone,
		NikeName: user.NikeName,
		RegDate:  user.RegDate.Format("2006-01-02"),
	}, nil
}

func Register(id string, pw string) *UserError {
	user := &models.UserPO{}

	user.Uuid = strings.Replace(uuid.NewV4().String(), "-", "", -1)
	if utils.RegCheck(id, utils.PHONE_REG, utils.MAIL_REG) == 1 {
		user.Phone = id
		user.Email = sql.NullString{}.String
	} else {
		user.Email = id
		user.Phone = sql.NullString{}.String
	}

	user.NikeName = "user-" + user.Uuid[26:]

	user.Password, user.Salt, _ = utils.EncodeWithRS(pw)

	user.RegDate = time.Now()

	logs.Info("register user: %v", user)

	if err := models.InsertUser(user); err != nil {
		logs.Info("register err: %v", err)
		return &UserError{RepeatId: true, Msg: "账户已存在"}
	}

	return nil
}
