package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Address struct {
	Uuid      string `json:"uuid"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Addressee string `json:"addressee"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	IsDefault bool   `json:"isDefault"`
}

// AddAddress 添加地址
func AddAddress(address *Address) (string, bool) {
	newOrm := orm.NewOrm()
	exec, err := newOrm.Raw(
		`
		insert into address (uuid, id, name, addressee, address, phone, is_default, ava) values (
		    ?,
		    ?,
		    ?,
		 	?,
		    ?,
			?,
		    (
		        case when
		        (
		            select a.num from
		                (select COUNT(1) as num from address where uuid = ?) as a
		        ) > 0
		        then 0
		        else 1
		        end
		    ),
		    (
		        case when
		        (
		            select a.num from
		                (select COUNT(1) as num from address where uuid = ?) as a
		        ) < 10
		        then 1
		        else 0
		        end
		    )
		);`,
		address.Uuid, address.Id, address.Name, address.Addressee, address.Address, address.Phone, address.Uuid, address.Uuid).Exec()

	if err != nil {
		logs.Error("insert address error: %v", err)
		return "", false
	}

	affected, err := exec.RowsAffected()

	logs.Info("affected line: %v, err: %v", affected, err)

	return address.Id, err == nil && affected == 1
}

func DeleteAddress(id string, uuid string) bool {

	exec, err := orm.NewOrm().Raw(
		`delete from address where id = ? and uuid = ? and is_default = 0;`,
		id, uuid).Exec()

	if err != nil {
		return false
	}

	affected, err := exec.RowsAffected()

	return err == nil && affected == 1
}

func SetDefault(id string, uuid string) bool {

	o := orm.NewOrm()

	err := o.Begin()

	_, err = o.Raw(
		`update address set is_default = 0 where uuid = ?`, uuid).Exec()

	exec, err := o.Raw(
		`update address set is_default = 1 where id = ? and uuid = ?;`,
		id, uuid).Exec()

	affected, err := exec.RowsAffected()

	if err != nil || (err != nil && affected != 1) {
		logs.Error("set default error: %v", err)
		_ = o.Rollback()
		return false
	}

	_ = o.Commit()

	return true
}

func GetAllAddresses(userId string) []Address {

	var addresses []Address

	num, err := orm.NewOrm().Raw(
		`select id, name, addressee, address, phone, is_default from address where uuid = ?`,
		userId).QueryRows(&addresses)

	logs.Info("query address, num = %v, err = %v", num, err)

	if err != nil {
		return nil
	}

	return addresses
}
