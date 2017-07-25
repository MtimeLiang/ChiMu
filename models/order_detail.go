package models

import (
	"github.com/astaxie/beego/orm"
)

type OrderDetail struct {
	ID          int `json:"id"`
	OID         int `json:"oid"`
	PID         int `json:"pid"`
	Count       int
	Price       float64
	ProductInfo Product
}

func AddOrderDetail(detail OrderDetail) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO orderdetail(oid, pid, count, price) VALUES(?, ?, ?, ?)",
		detail.OID, detail.PID, detail.Count, detail.Price).Exec()
}

func GetOrderDetailByOID(OID int) OrderDetail {
	var detail OrderDetail
	o := orm.NewOrm()
	o.Raw("SELECT * FROM orderdetail WHERE oid = ?", OID).QueryRow(&detail)
	return detail
}
