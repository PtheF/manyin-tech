package routers

import (
	"github.com/astaxie/beego"
	"go-service/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{}, "GET:Index")
	//
	//beego.Router("/login", &controllers.MainController{}, "GET:Login")
	//
	//beego.Router("/register", &controllers.MainController{}, "GET:Register")
	//
	//beego.Router("/home", &controllers.MainController{}, "GET:Home")

	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/login", &controllers.UserController{}, "POST:Login"),
		beego.NSRouter("/register", &controllers.UserController{}, "POST:Register"),
		beego.NSRouter("/captcha", &controllers.UserController{}, "GET:Captcha"),
		beego.NSRouter("/space", &controllers.SpaceController{}, "GET:Space"),

		beego.NSRouter("/address", &controllers.AddressController{}, "PUT:AddAddress"),
		beego.NSRouter("/address/:id", &controllers.AddressController{}, "POST:SetDefault"),
		beego.NSRouter("/address/:id", &controllers.AddressController{}, "DELETE:DelAddress"),

		beego.NSRouter("/enterprise", &controllers.EnterpriseController{}, "PUT:AddEnterprise"),

		beego.NSNamespace("/shop",
			beego.NSRouter("/types", &controllers.ProductController{}, "GET:GetProductTypes"),
			beego.NSRouter("/list/:typeId([0-9]+)/:page([0-9]+)", &controllers.ProductController{}, "GET:GetProdsByTypeId"),
			beego.NSRouter("/item/:prodId", &controllers.ProductController{}, "GET:GetProductDetailById"),
		),

		beego.NSNamespace("/account",
			beego.NSRouter("/verify", &controllers.AccountController{}, "POST:VerifyAccount"),
			beego.NSRouter("/info", &controllers.AccountController{}, "POST:UpdateAccount"),
			beego.NSRouter("/phone", &controllers.AccountController{}, "POST:UpdatePhone"),
			beego.NSRouter("/email", &controllers.AccountController{}, "POST:UpdateEmail"),
			beego.NSRouter("/pw", &controllers.AccountController{}, "POST:UpdatePassword"),
		),

		beego.NSNamespace("/order",
			beego.NSRouter("/token", &controllers.OrderController{}, "GET:NextOrderToken"),
			beego.NSRouter("/create/:token", &controllers.OrderController{}, "PUT:CreateOrder"),
			beego.NSRouter("/status/:orderId([0-9]+)", &controllers.OrderController{}, "GET:GetOrderStatus"),
			beego.NSRouter("/list", &controllers.OrderController{}, "GET:GetOrderList"),
		),

		beego.NSNamespace("/pay",
			beego.NSRouter("/orderInfo/:orderId([0-9]+)", &controllers.OrderController{}, "GET:GetUnpaidOrderInfo"),
			beego.NSRouter("/:orderId([0-9]+)", &controllers.OrderController{}, "POST:PayOrder"),
		),

		beego.NSNamespace("/admin",
			beego.NSRouter("/login", &controllers.AdminController{}, "POST:Login"),
			beego.NSRouter("/product", &controllers.AdminController{}, "PUT:AddProduct"),
			beego.NSRouter("/product/:prodId", &controllers.AdminController{}, "DELETE:DelProduct"),
			beego.NSRouter("/types", &controllers.AdminController{}, "POST:UpdateProductType"),
		),
	)

	beego.AddNamespace(ns)
}
