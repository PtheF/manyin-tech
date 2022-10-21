package controllers

import (
	"github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego/logs"
	"go-service/models"
	"go-service/service"
	"go-service/utils"
)

type AdminController struct {
	BaseController
}

const (
	AdminPassword    = "xxxxx"
	AdminTokenPrefix = "admin:login:"
)

func (c *AdminController) Login() {
	pw := c.GetString("pw")

	if pw == AdminPassword {
		uuid := utils.NextUuid()

		_ = template.OpsForValue().SetEx(AdminTokenPrefix+uuid, "1", 60*60*2)

		c.Success(uuid)
	}

	c.Failed("登录失败")
}

func (c *AdminController) DelProduct() {
	prodId := c.RestString("prodId")

	if err := service.DelProduct(prodId); err != nil {
		c.Failed(err.Error())
	}

	c.Success("删除成功")
}

func (c *AdminController) UpdateProductType() {
	var prodType models.ProductType

	c.GetJson(&prodType)

	logs.Info("update product type: %v", prodType)

	if err := service.UpdateProdType(prodType); err != nil {
		c.Failed(err.Error())
	}

	c.Success("修改成功")
}

func (c *AdminController) AddProduct() {

	var prod models.Product

	c.GetJson(&prod)

	//logs.Info("prod:")
	//logs.Info(prod)

	if err := service.AddProduct(&prod); err != nil {
		logs.Error("AdminController > service.AddProduct error: %v", err)
		c.Failed(err.Error())
	}

	c.Success("添加成功")

	//prodAvatar := c.GetString("prodAvatar")
	//
	//logs.Info("img:")
	//logs.Info(prodAvatar)
	//
	//err := utils.SaveBase64Img("test123", "/usr/local/www/assets/img/", prodAvatar)
	//
	//if err != nil {
	//	logs.Error(err)
	//	c.Failed("save err")
	//}

	//c.Failed("error")

	//img, _ := base64.StdEncoding.DecodeString(prodAvatar)
	//
	//err := ioutil.WriteFile("/usr/local/www/assets/test.png", img, 0666)
	//
	//if err != nil {
	//	logs.Error(err)
	//	c.Failed("save error")
	//}

	//c.Success("save ok")

}
