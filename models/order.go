package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Order struct {
	ID         int    `json:"id"`
	UID        int    `json:"uid"`
	OrderNum   string `json:"order_num"`
	Status     int
	Count      int
	Memo       string
	AddressID  int `json:"address_id"`
	CouponID   int `json:"coupon_id"`
	Amount     float64
	Pay        float64
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}

func AddOrder(order *Order) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO `order`(uid, order_num, status, count, amount, create_time, modify_time) VALUES(?, ?, ?, ?, ?, ?, ?)",
		order.UID, order.OrderNum, order.Status, order.Count, order.Amount, order.CreateTime, order.ModifyTime).Exec()
}

func GetOrderList() []Order {
	var orders []Order
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `order`").QueryRows(&orders)
	return orders
}

func GetOrderByID(ID int) Order {
	var order Order
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `order` WHERE id = ?", ID).QueryRow(&order)
	return order
}

func GetOrderByUID(UID int) Order {
	var order Order
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `order` WHERE uid = ?", UID).QueryRow(&order)
	return order
}

func GetOrderByOrderNum(orderNum string) Order {
	var order Order
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `order` WHERE order_num = ?", orderNum).QueryRow(&order)
	return order
}

func ModifyOrder(order Order) {
	o := orm.NewOrm()
	o.Raw("UPDATE `order` SET address_id = ?, coupon_id = ?, status = ?, memo = ?, pay = ?, modify_time = ? WHERE id = ?",
		order.AddressID, order.CouponID, order.Status, order.Memo, order.Pay, order.ModifyTime, order.ID).Exec()
}
