package test

import (
	"github.com/astaxie/beego/logs"
	"go-service/models"
	"testing"
)

func TestGetProdPage(t *testing.T) {
	prodPage, err := models.GetProductByType(2, 12, 3)

	if err != nil {
		logs.Error("get prod by type error: %v", err)
	}

	logs.Info("total page = %v", prodPage.TotalPage)

	logs.Info("prod list len = %v", len(prodPage.Products))

	for prodVo := range prodPage.Products {
		logs.Info(prodPage.Products[prodVo])
	}
}
