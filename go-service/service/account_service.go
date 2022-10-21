package service

import (
	"errors"
	"github.com/PtheF/go-redis-template/template"
	"go-service/models"
	"go-service/utils"
)

const (
	UpdateVerify = "update:"
)

func VerifyService(captcha string, pw string, tp string, user *models.UserDTO) error {
	if tp == "pw" {
		userPO, err := models.GetUserByUuid(user.Uuid)

		if err != nil {
			return err
		}

		if pw, err := utils.EncodeWithSalt(pw, userPO.Salt); err != nil || pw != userPO.Password {
			return errors.New("认证不通过")
		}
	} else {
		if _, err := NewCaptcha(user.Phone).Verify(captcha); err != nil {
			if _, err = NewCaptcha(user.Email).Verify(captcha); err != nil {
				return err
			}
		}
	}

	if err := template.OpsForValue().SetEx(UpdateVerify+user.Uuid, "1", 600); err != nil {
		return err
	}

	return nil
}

func UpdateAccount(user *models.UserDTO) error {

	return models.UpdateAccount(user)
}

func UpdateAccountWithVerify(cpKey string, cp string, user *models.UserDTO) error {

	if b, _ := template.OpsForKey().Exist(UpdateVerify + user.Uuid); !b {
		return errors.New("未在10分钟内完成修改")
	}

	if _, err := NewCaptcha(cpKey).Verify(cp); err != nil {
		return err
	}

	if err := models.UpdateAccount(user); err != nil {
		return err
	}

	return nil
}

func UpdatePassword(pw string, userId string) error {

	pw, salt, _ := utils.EncodeWithRS(pw)

	if err := models.UpdatePassword(pw, salt, userId); err != nil {
		return err
	}

	return nil
}
