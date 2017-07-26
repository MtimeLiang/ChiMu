package models

import (
	"github.com/astaxie/beego/orm"
)

type OrderDetail struct {
	ID          int     `json:"id" orm:"column(id)"`
	OID         int     `json:"oid" orm:"column(oid)"`
	PID         int     `json:"pid" orm:"column(pid)"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
	ProductInfo Product `json:"product_info"`
}

func AddOrderDetail(detail *OrderDetail) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO orderdetail(oid, pid, count, price) VALUES(?, ?, ?, ?)",
		detail.OID, detail.PID, detail.Count, detail.Price).Exec()
}

func GetOrderDetailByOID(OID int) []OrderDetail {
	var details []OrderDetail
	o := orm.NewOrm()
	o.Raw("SELECT * FROM orderdetail WHERE oid = ?", OID).QueryRows(&details)
	return details
}
