package controllers

//
//import "go-service/service"
//
//type MainController struct {
//	BaseController
//}
//
//func (c *MainController) Index() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}
//
//func (c *MainController) Login() {
//	c.TplName = "login.html"
//	c.TplExt = "html"
//
//}
//
//func (c *MainController) Register() {
//	c.TplName = "register_temp.html"
//	c.TplExt = "html"
//}
//
//func (c *MainController) Home() {
//	c.TplName = "home.html"
//	c.TplExt = "html"
//	//c.EnableRender = false
//
//	user := c.GetUser()
//
//	addresses := service.GetAllAddress(user.Uuid)
//
//	c.Data["addresses"] = addresses
//	c.Data["addresses_len"] = len(addresses)
//	c.Data["user"] = user
//}
