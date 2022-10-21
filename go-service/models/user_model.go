package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserDTO struct {
	Uuid     string `json:"uuid"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	NikeName string `json:"nikeName"`
	RegDate  string `json:"regDate"`
}

type UserPO struct {
	Uuid     string
	Email    string
	Phone    string
	Password string
	NikeName string
	Salt     string
	RegDate  time.Time
}

func empty2Nil(s string) interface{} {
	if s == "" {
		return nil
	}

	return s
}

func InsertUser(u *UserPO) error {
	o := orm.NewOrm()

	exec, err := o.Raw(`
			insert into user 
			    (uuid, nike_name, email, phone, salt, reg_date, password) 
			values (?, ?, ?, ?, ?, ?, ?)`,
		u.Uuid, u.NikeName, empty2Nil(u.Email), empty2Nil(u.Phone), u.Salt, u.RegDate, u.Password).Exec()

	if err != nil {
		return err
	}

	if affected, err := exec.RowsAffected(); err != nil || affected != 1 {
		return errors.New(fmt.Sprintf("AffectedNumberError, affected=%d, err=%v", affected, err))
	}

	return nil
}

func GetUserByEmail(email string) (user *UserPO, err error) {

	user = &UserPO{}

	o := orm.NewOrm()
	err = o.Raw("select uuid, email, phone, nike_name, password, salt, reg_date from user where email = ?", email).QueryRow(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByPhone(phone string) (user *UserPO, err error) {
	user = &UserPO{}

	o := orm.NewOrm()
	err = o.Raw("select uuid, email, phone, nike_name, password, salt, reg_date from user where phone = ?", phone).QueryRow(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUuid(id string) (user *UserPO, err error) {
	user = &UserPO{}

	err = orm.NewOrm().Raw(
		`select uuid, email, phone, nike_name, password, salt, reg_date from user where uuid = ?`,
		id).QueryRow(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateAccount(user *UserDTO) error {
	exec, err := orm.NewOrm().Raw(
		`
			update user set
			email = ?, phone = ?, nike_name = ?
			where uuid = ?
		`,
		empty2Nil(user.Email), empty2Nil(user.Phone), user.NikeName, user.Uuid).Exec()

	if err != nil {
		return err
	}

	if affected, err := exec.RowsAffected(); err != nil || affected != 1 {
		return errors.New(fmt.Sprintf("AffectedNumberError, affected=%d, err=%v", affected, err))
	}

	return nil
}

func UpdatePassword(pw string, salt string, uid string) error {

	exec, err := orm.NewOrm().Raw(`
		update user set password = ?, salt = ? where uuid = ? 
	`, pw, salt, uid).Exec()

	if err != nil {
		return err
	}

	if affected, err := exec.RowsAffected(); err != nil || affected != 1 {
		return errors.New(fmt.Sprintf("AffectedNumberError, affected=%d, err=%v", affected, err))
	}

	return nil
}
