package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

/**
create table order (
	order_id int(64) primary key,
	user_id   ,
	prod_id   ,

	prod_name    ,
	prod_avatar   ,

	total_price int(64) not null,
	buy_num int not null,

	addressee  ,
	address  ,
	addressee_phone  ,

	create_time  ,
	status  int not null,
);
*/

type OrderPO struct {
	OrderId    int64  `json:"orderId"`
	OrderIdStr string `json:"orderIdStr"`
	UserId     string `json:"userId"`
	ProdId     string `json:"prodId"`

	ProdName   string `json:"prodName"`
	ProdAvatar string `json:"prodAvatar"`

	TotalPrice int64 `json:"totalPrice"`
	BuyNum     int   `json:"buyNum"`

	Addressee      string `json:"addressee"`
	Address        string `json:"address"`
	AddresseePhone string `json:"addresseePhone"`

	CreateTime time.Time `json:"createTime"`
	Status     int8      `json:"status"`
}

func AddOrder(order *OrderPO) error {

	exec, err := orm.NewOrm().Raw(`
		insert into paid_order values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`,
		order.OrderId,
		order.UserId,
		order.ProdId,
		order.ProdName,
		order.ProdAvatar,
		order.TotalPrice,
		order.BuyNum,
		order.Addressee,
		order.Address,
		order.AddresseePhone,
		order.CreateTime,
		order.Status).Exec()

	if err != nil {
		logs.Error("insert paid order error: %v", err)
		return err
	}

	affected, _ := exec.RowsAffected()

	if affected != 1 {
		return errors.New("affect != 1")
	}

	return nil
}

func GetOrdersByUserId(userId string) ([]OrderPO, error) {

	var orders []OrderPO

	if _, err := orm.NewOrm().Raw(
		`select * from paid_order where user_id = ?`, userId).QueryRows(&orders); err != nil {
		logs.Error("query paid order error: user=%v, err=%v", userId, err)
		return nil, err
	}

	return orders, nil
}
