package models

import (
	"github.com/astaxie/beego/orm"
)

type Cart struct {
	ID         int `json:"id"`
	UID        int `json:"uid"`
	PID        int `json:"pid"`
	Count      int
	Title      string
	SubMessage string `json:"submessage"`
	Price      float64
	Volume     int
	Images     string
}

func AddCart(c *Cart) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO cart(pid, uid, count) VALUES (?, ?, ?)", c.PID, c.UID, c.Count).Exec()
}

func GetCart(PID, UID int) Cart {
	var cart Cart
	o := orm.NewOrm()
	o.Raw("SELECT * FROM cart WHERE pid = ? AND uid = ?", PID, UID).QueryRow(&cart)
	return cart
}

func ModifyCartByID(c *Cart) {
	o := orm.NewOrm()
	o.Raw("UPDATE cart SET count = ? WHERE id = ?", c.Count, c.ID).Exec()
}

func GetCartByUID(UID int) []Cart {
	var carts []Cart
	o := orm.NewOrm()
	o.Raw("SELECT * FROM cart WHERE uid = ?", UID).QueryRows(&carts)
	return carts
}

func DeleteCart(ID int) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM cart WHERE id = ?", ID).Exec()
}
