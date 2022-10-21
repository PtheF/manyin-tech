package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-service/models"
	"go-service/service"
	"time"
)

type EnterpriseController struct {
	BaseController
}

func (c *EnterpriseController) AddEnterprise() {
	//enterprise := models.EnterprisePO{}
	var enterprise models.EnterprisePO

	c.GetJson(&enterprise)

	enterprise.UserId = c.GetUser().Uuid
	enterprise.CreateDate = time.Now()

	logs.Info("enterprise info: %v", enterprise)

	if err := service.AddEnterprise(&enterprise); err != nil {
		c.Failed(err.Error())
	} else {
		c.Success("信息录入成功，等待审核")
	}
}
