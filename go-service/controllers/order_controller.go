package controllers

import (
	"github.com/astaxie/beego/logs"
	"go-service/models"
	"go-service/service"
	"strconv"
)

type OrderController struct {
	BaseController
}

func (c *OrderController) CreateOrder() {
	//order := models.OrderPO{}

	var order models.OrderPO

	c.GetJson(&order)

	orderToken := c.RestString("token")

	order.UserId = c.GetUser().Uuid

	if orderId, err := service.CreateOrder(&order, orderToken); err != nil {
		c.Failed(err.Error())
	} else {

		logs.Warning("gen order id: %v", orderId)

		//c.ReturnData(orderId)

		// 这里注意下：不知道为啥 go 的 json 发送 int64 会出现精度丢失的情况，
		// 这里出现了订单 id 发送到前端发生变化的情况，所以这里改用字符串发送
		c.ReturnData(strconv.FormatInt(orderId, 10))
	}
}

func (c *OrderController) NextOrderToken() {
	if token, err := service.NextOrderToken(); err != nil {
		c.Failed(err.Error())
	} else {
		c.Success(token)
	}
}

func (c *OrderController) GetOrderStatus() {
	orderId := c.RestString("orderId")

	logs.Info("check order status: orderId=%v", orderId)

	if status, err := service.GetOrderStatus(orderId); err != nil {
		c.Failed("未找到待支付订单，支付已过期")
	} else {
		c.ReturnData(status)
	}
}

func (c *OrderController) GetOrderList() {
	userId := c.GetUser().Uuid

	if orders, err := service.GetUserOrders(userId); err != nil {
		c.Failed(err.Error())
	} else {
		c.ReturnData(orders)
	}
}

func (c *OrderController) GetUnpaidOrderInfo() {
	orderId := c.RestString("orderId")

	if order, err := service.GetUnpaidOrder(orderId); err != nil {
		logs.Error("get order info error: %v", err)
		c.Failed(err.Error())
	} else {
		c.ReturnData(order)
	}
}

func (c *OrderController) PayOrder() {
	orderId := c.RestString("orderId")

	if err := service.PayOrder(orderId); err != nil {
		logs.Error("pay order error: %v", err)
		c.Failed(err.Error())
	}

	c.Success("付款成功")
}
