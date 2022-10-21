package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ProductType struct {
	ProdTId     int    `json:"prodTypeId"`
	TName       string `json:"tName"`
	SubText     string `json:"subText"`
	Description string `json:"description"`
}

func GetProductTypes() (prodTypes []ProductType) {

	if _, err := orm.NewOrm().Raw(`
		select prod_t_id, t_name, sub_text, description from product_types 
			`).QueryRows(&prodTypes); err != nil {
		logs.Error("GetProductTypesError: %v", err)
		return nil
	}

	return prodTypes
}

func UpdateProductType(prodType ProductType) error {
	exec, err := orm.NewOrm().Raw(
		`update product_types set t_name = ?, sub_text = ?, description = ? where prod_t_id = ?`,
		prodType.TName, prodType.SubText, prodType.Description, prodType.ProdTId,
	).Exec()

	if err != nil {
		return err
	}

	if affected, _ := exec.RowsAffected(); affected != 1 {
		return errors.New("修改失败，您压根就没改")
	}

	return nil
}
