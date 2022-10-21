package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Product struct {
	ProdId      string `json:"prodId"`
	ProdName    string `json:"prodName"`
	ProdDesc    string `json:"prodDesc"`
	ProdTypeId  int    `json:"prodTypeId"`
	ProdType    string `json:"prodType"`
	ProdAvatar  string `json:"prodAvatar"`
	ProdDigital string `json:"prodDigital"`
	ProdPrice   int64  `json:"prodPrice"`
	ProdUnit    string `json:"prodUnit"`
}

type ProductListVO struct {
	ProdId     string `json:"prodId"`
	ProdName   string `json:"prodName"`
	ProdPrice  int64  `json:"prodPrice"`
	ProdUnit   string `json:"prodUnit"`
	ProdAvatar string `json:"prodAvatar"`
}

type ProductPage struct {
	Products    []ProductListVO `json:"products"`
	CurrentPage int             `json:"currentPage"`
	TotalPage   int             `json:"totalPage"`
	PageCount   int             `json:"pageCount"`
	TypeId      int             `json:"typeId"`
}

/*

create table product (
    prod_id char(40) primary key,
    prod_name varchar(30) not null,
	prod_desc text not null,
	prod_type_id tinyint not null,
	prod_avatar text not null,
	prod_digital text not null,
	prod_price int(64) not null,
	prod_unit char(5) not null,
    foreign key (prod_type_id) references product_types(prod_t_id)
);

*/

func AddProduct(prod *Product) error {

	//logs.Info("get prod: id=%v, name=%v, desc=%v, type")

	exec, err := orm.NewOrm().Raw(
		`insert into product values (?, ?, ?, ?, ?, ?, ?, ?)`,
		prod.ProdId, prod.ProdName, prod.ProdDesc, prod.ProdTypeId, prod.ProdAvatar, prod.ProdDigital, prod.ProdPrice, prod.ProdUnit,
	).Exec()

	if err != nil {
		logs.Error("exec error: %v", err)
		return err
	}

	affected, err := exec.RowsAffected()

	if err != nil {
		logs.Error("add product error:%v", err)
		return err
	}

	if affected != 1 {
		return errors.New("affected != 1")
	}

	return nil
}

func GetProductByType(currentPage int, pageCount int, typeId int) (*ProductPage, error) {

	o := orm.NewOrm()

	var prodCount int

	err := o.Raw("select COUNT(1) from product where prod_type_id = ?", typeId).QueryRow(&prodCount)

	//logs.Info("prodCount=%v", prodCount)

	var totalPage int

	if (prodCount % pageCount) == 0 {
		totalPage = prodCount / pageCount
	} else {
		totalPage = (prodCount / pageCount) + 1
	}

	var prodList []ProductListVO

	_, err = o.Raw(
		`select prod_id, prod_name, prod_price, prod_unit, prod_avatar from product where prod_type_id = ? limit ? offset ?`,
		typeId, pageCount, (currentPage-1)*pageCount,
	).QueryRows(&prodList)

	if err != nil {
		logs.Error("get product list error: %v", err)
		return nil, err
	}

	productPage := &ProductPage{
		Products:    prodList,
		PageCount:   pageCount,
		TotalPage:   totalPage,
		TypeId:      typeId,
		CurrentPage: currentPage,
	}

	return productPage, nil
}

func GetProductById(prodId string) (Product, error) {

	var prod Product

	if err := orm.NewOrm().Raw(
		`select * from product where prod_id = ?`,
		prodId).QueryRow(&prod); err != nil {
		logs.Error("get product error: %v", err)
		return prod, err
	}

	logs.Info("get product by id models: %v", prod)

	return prod, nil
}

func DelProduct(prodId string) error {
	exec, err := orm.NewOrm().Raw(
		`delete from product where prod_id = ?`,
		prodId).Exec()

	if err != nil {
		return err
	}

	if affected, err := exec.RowsAffected(); affected != 1 || err != nil {
		return errors.New("错误，影响行数 != 0")
	}

	return nil
}
