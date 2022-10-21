package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type EnterprisePO struct {
	UserId        string    `json:"userId"`
	FullName      string    `json:"fullName"`
	SocialCode    string    `json:"socialCode"`
	RegistryMoney int64     `json:"registryMoney"`
	EnterAddress  string    `json:"enterAddress"`
	OperatorName  string    `json:"operatorName"`
	OperatorEmail string    `json:"operatorEmail"`
	OperatorPhone string    `json:"operatorPhone"`
	CreateDate    time.Time `json:"createDate"`
	Status        int       `json:"status"`
	Captcha       string    `json:"captcha"`
}

//type EnterpriseDTO struct {
//	FullName      string
//	SocialCode    string
//	RegistryMoney int64
//	EnterAddress  string
//	OperatorName  string
//	OperatorEmail string
//	OperatorPhone string
//	CreateDate    time.Time
//	Status        int
//}

func AddEnterprise(e *EnterprisePO) error {

	exec, err := orm.NewOrm().Raw(`
		insert into enterprise_info
		(
		 user_id,
		 full_name,
		 social_code,
		 registry_money,
		 enter_address,
		 operator_name,
		 operator_email,
		 operator_phone,
		 create_date,
		 status
		)values (
		    ?, ?, ?, ?, ?, ?, ?, ?, ?, 1
		);
	`, e.UserId, e.FullName, e.SocialCode, e.RegistryMoney, e.EnterAddress, e.OperatorName,
		e.OperatorEmail, e.OperatorPhone, e.CreateDate).Exec()

	if err != nil {
		logs.Error("add enterprise info err: %v", err)
		return err
	}

	if affected, err := exec.RowsAffected(); err != nil || affected != 1 {
		if err != nil {
			return err
		}

		return errors.New(fmt.Sprintf("add error, affected row = %d", affected))
	}

	return nil
}

func GetEnterpriseByUser(userId string) *EnterprisePO {

	enterprise := &EnterprisePO{}

	err := orm.NewOrm().Raw(`
		select user_id,
		 full_name,
		 social_code,
		 registry_money,
		 enter_address,
		 operator_name,
		 operator_email,
		 operator_phone,
		 create_date,
		 status from enterprise_info where user_id = ?
	`, userId).QueryRow(enterprise)

	if err != nil {
		return nil
	}

	return enterprise
}
