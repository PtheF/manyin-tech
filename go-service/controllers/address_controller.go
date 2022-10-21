package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-service/models"
	"go-service/service"
)

type AddressController struct {
	BaseController
}

// AddAddress 添加地址，PUT FORM address=...
func (c *AddressController) AddAddress() {

	address := &models.Address{}

	c.GetJson(address)

	if len(address.Name) >= 20 ||
		len(address.Addressee) >= 30 ||
		len(address.Address) >= 255 ||
		len(address.Phone) > 20 {
		c.Resp(ParamIllegal)
	}

	address.Uuid = c.GetUser().Uuid

	logs.Info("get address: %v", address)

	if id, ok := service.AddAddress(address); !ok {
		c.Failed("添加失败")
	} else {
		c.ReturnData(id)
	}
}

func (c *AddressController) DelAddress() {
	aid := c.RestString("id")
	userId := c.GetUser().Uuid

	if service.DeleteAddress(aid, userId) {
		c.Success("删除成功")
	} else {
		c.Failed("删除失败")
	}
}

func (c *AddressController) SetDefault() {
	aid := c.RestString("id")
	userId := c.GetUser().Uuid

	if service.SetDefault(aid, userId) {
		c.Success("修改默认地址成功")
	} else {
		c.Failed("设置失败,鬼知道你在干啥")
	}
}
