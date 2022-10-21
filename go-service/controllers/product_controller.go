package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-service/service"
	"strconv"
)

type ProductController struct {
	BaseController
}

func (c *ProductController) GetProductTypes() {
	if prodTypes := service.GetProductTypes(); prodTypes == nil {
		c.Failed("NoProductTypes")
	} else {
		c.ReturnData(prodTypes)
	}
}

func (c *ProductController) GetProdsByTypeId() {

	logs.Info("get prod by type id")

	typeId, err := strconv.Atoi(c.RestString("typeId"))
	page, err := strconv.Atoi(c.RestString("page"))

	if err != nil {
		c.Resp(ParamIllegal)
	}

	logs.Info("Get Product List Type Id = %v, Page = %v;", typeId, page)

	prodPage, err := service.GetProductPageByType(page, 12, typeId)

	if err != nil {
		c.Failed(err.Error())
	}

	c.ReturnData(prodPage)
}

func (c *ProductController) GetProductDetailById() {
	prodId := c.RestString("prodId")

	prod, err := service.GetProductById(prodId)

	if err != nil {
		logs.Error("ProductController > service.GetProductById error: %v", err)
		c.Failed(err.Error())
	}

	c.ReturnData(prod)
}
