package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-service/models"
	"go-service/service"
)

type SpaceData struct {
	Addresses  []models.Address    `json:"addresses"`
	Enterprise models.EnterprisePO `json:"enterprise"`
}

type SpaceController struct {
	BaseController
}

func (c *SpaceController) Space() {
	userId := c.GetUser().Uuid

	addresses := service.GetAllAddress(userId)
	enterprise := service.GetEnterpriseByUser(userId)

	spaceData := SpaceData{}

	if enterprise != nil {
		logs.Info("get enterprise info")
		logs.Info(enterprise)
		spaceData.Enterprise = *enterprise
	}
	spaceData.Addresses = addresses

	c.ReturnData(spaceData)
}
