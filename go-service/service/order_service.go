package service

import (
	"encoding/json"
	"errors"
	"github.com/PtheF/go-redis-template/template"
	"github.com/astaxie/beego/logs"
	"go-service/models"
	"go-service/utils"
	"io/ioutil"
	"strconv"
	"time"
)

/**
create a temp product
prodId: 111
stock: 1000
*/

const (
	CreateOrderToken = "{order}:create:token:"
	OrderInfo        = "{order}:id:"
	UserUnpaidOrders = "{order}:user:"
)

var (
	CreateOrderScript, _ = ioutil.ReadFile("./lua/create_order.lua")
	PayOrderScript, _    = ioutil.ReadFile("./lua/pay_order.lua")
)

func NextOrderToken() (string, error) {
	token := utils.NextUuid()
	if err := template.OpsForValue().SetEx(CreateOrderToken+token, "1", 60*10); err != nil {
		logs.Error("redis setEx error: %v", err)
		return "", err
	}

	return token, nil
}

// GetOrderStatus
// return 1 = payed
// return -1 = order not found
// return 0 = no pay
func GetOrderStatus(orderId string) (int, error) {

	logs.Info("order status key: %v", OrderInfo+orderId)

	if status, err := template.OpsForHash().Get(OrderInfo+orderId, "status"); err != nil {
		logs.Error("redis error: %v", err)
		return -1, err
	} else {
		if atoi, err := strconv.Atoi(status); err != nil {
			logs.Error("Atoi error: %v", err)
			return -1, err
		} else {
			return atoi, nil
		}
	}
}

func GetUnpaidOrder(orderId string) (*models.OrderPO, error) {
	if get, err := template.OpsForHash().Get(OrderInfo+orderId, "status"); err != nil || get != "0" {
		return nil, errors.New("支付完成或过期")
	}

	if get, err := template.OpsForHash().Get(OrderInfo+orderId, "info"); err != nil {
		logs.Error("get unpaid info error: %v", err)
		return nil, err
	} else {

		logs.Info("get: %v", get)

		var orderInfo models.OrderPO
		err = json.Unmarshal([]byte(get), &orderInfo)

		return &orderInfo, err
	}
}

func GetUserOrders(userId string) ([]models.OrderPO, error) {

	logs.Info("get orders user id:%v", userId)

	//orders := []models.OrderPO{}
	orders := make([]models.OrderPO, 0)

	if unpaidOrdersId, err := template.OpsForSet().Members(UserUnpaidOrders + userId); err == nil {
		for i := 0; i < len(unpaidOrdersId); i++ {

			logs.Info("each unpaid order id: unpaid order[%v] id = %v", i, unpaidOrdersId[i])

			if unpaidOrderInfo, err := template.OpsForHash().Get(OrderInfo+unpaidOrdersId[i], "info"); err == nil {
				var order models.OrderPO
				if err = json.Unmarshal([]byte(unpaidOrderInfo), &order); err == nil {

					logs.Info("get unpaid order info: %v", order)

					orders = append(orders, order)
				} else {
					logs.Error("json unmarshal unpaid order info error: %v", err)
				}
			} else {
				logs.Error("redis get unpaid order info error: %v", err)
			}
		}
	} else {
		logs.Error("get all unpaid orders error: %v", err)
	}

	//logs.Info("--------------------------------------")
	//
	//for i := 0; i < len(orders); i++ {
	//	logs.Info("orderInfo: id=%v, name=%v", orders[i].OrderId, orders[i].ProdName)
	//}

	if paidOrders, err := models.GetOrdersByUserId(userId); err != nil {
		logs.Warning("paid order empty: err=%v", err)
	} else {
		orders = append(orders, paidOrders...)
	}

	for i := 0; i < len(orders); i++ {
		orders[i].OrderIdStr = strconv.FormatInt(orders[i].OrderId, 10)
	}

	return orders, nil
}

func PayOrder(orderId string) error {

	execute, err := template.Execute(string(PayOrderScript), []interface{}{"{order}"}, []interface{}{orderId})

	if err != nil {
		return err
	}

	exec := string(execute.([]uint8))

	logs.Info("pay order script exec res: %v", execute)

	if exec == "nil" {
		return errors.New("订单已支付或订单过期")
	}

	var order models.OrderPO

	err = json.Unmarshal([]byte(exec), &order)

	if err != nil {
		return err
	}

	order.Status = 1

	logs.Info("order info: %v", order)

	return models.AddOrder(&order)
}

func CreateOrder(order *models.OrderPO, orderToken string) (int64, error) {
	order.OrderId = utils.NextId("order")

	//orderToken := utils.NextUuid()

	//payToken := utils.NextUuid()

	order.CreateTime = time.Now()

	orderInfoJsonStr, _ := json.Marshal(order)

	execute, err := template.Execute(
		string(CreateOrderScript),
		[]interface{}{"{order}"},
		[]interface{}{order.BuyNum, order.ProdId, order.OrderId, orderToken, order.UserId, string(orderInfoJsonStr)},
	)

	if err != nil {
		logs.Error("execute lua script error: %v", err)
	}

	logs.Info("create order: id=%v, token=%v, script=%v", order.OrderId, orderToken, execute)

	switch execute.(int64) {
	case 1:

		return order.OrderId, nil

	case 2:

		return -1, errors.New("商品库存不足")

	case 3:

		return -1, errors.New("创建失败，请退出此页面重新购买")

	default:

		return -1, errors.New("unknown")
	}
}
