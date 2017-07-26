package models

import (
	"github.com/astaxie/beego/orm"
)

type Coupon struct {
	ID        int    `json:"id" orm:"column(id)"`
	PID       int    `json:"pid" orm:"column(pid)"`
	UID       int    `json:"uid" orm:"column(uid)"`
	Price     int    `json:"price"`
	MaxPrice  int    `json:"max_price"`
	Title     string `json:"title"`
	BuildTime string `json:"build_time"`
	EndTime   string `json:"end_time"`
}

func GetCouponByID(ID int) Coupon {
	var coupon Coupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM coupon WHERE id = ?", ID).QueryRow(&coupon)
	return coupon
}

func GetCouponByPID(PID int) []Coupon {
	var coupons []Coupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM coupon WHERE pid = ?", PID).QueryRows(&coupons)
	return coupons
}

func GetCouponByUID(UID int) []Coupon {
	var coupons []Coupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM coupon WHERE uid = ?", UID).QueryRows(&coupons)
	return coupons
}

func NumberOfCouponByUID(UID int) int {
	count := 0
	o := orm.NewOrm()
	// warning：这里取count值，函数调用方式待确认
	o.Raw("SELECT count(id) FROM coupon WHERE uid = ?", UID).QueryRow(&count)
	return count
}
