package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Order struct {
	ID           int           `json:"id" orm:"column(id)"`
	UID          int           `json:"uid" orm:"column(uid)"`
	OrderNum     string        `json:"order_num"`
	Status       int           `json:"status"`
	Count        int           `json:"count"`
	Memo         string        `json:"memo"`
	AddressID    int           `json:"address_id" orm:"column(address_id)"`
	CouponID     int           `json:"coupon_id" orm:"column(coupon_id)"`
	Amount       float64       `json:"amount"`
	Pay          float64       `json:"pay"`
	CreateTime   time.Time     `json:"create_time"`
	ModifyTime   time.Time     `json:"modify_time"`
	OrderDetails []OrderDetail `json:"order_details"`
	AddressInfo  Address       `json:"address_info"`
	CouponInfo   Coupon        `json:"coupon_info"`
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

func GetOrderByUID(UID int) []Order {
	var orders []Order
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `order` WHERE uid = ?", UID).QueryRows(&orders)
	return orders
}

func GetOrderByOrderNum(orderNum string) Order {
	var order Order
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `order` WHERE order_num = ?", orderNum).QueryRow(&order)
	return order
}

func ModifyOrder(order *Order) {
	o := orm.NewOrm()
	o.Raw("UPDATE `order` SET address_id = ?, coupon_id = ?, status = ?, memo = ?, pay = ?, modify_time = ? WHERE id = ?",
		order.AddressID, order.CouponID, order.Status, order.Memo, order.Pay, order.ModifyTime, order.ID).Exec()
}

func GetOrderByUIDAndStatus(UID, status int) []Order {
	o := orm.NewOrm()
	var orders []Order
	o.Raw("SELECT * FROM `order` WHERE uid = ? AND status = ? ORDER BY create_time DESC", UID, status).QueryRows(&orders)
	return orders
}
