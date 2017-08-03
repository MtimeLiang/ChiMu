package models

import "github.com/astaxie/beego/orm"

type MyCoupon struct {
	ID         int    `json:"id" orm:"column(id)"`
	UID        int    `json:"uid" orm:"column(uid)"`
	CouponID   int    `json:"coupon_id" orm:"column(coupon_id)"`
	Status     int    `json:"status"`
	CouponInfo Coupon `json:"couponInfo"`
}

func AddMyCoupon(m *MyCoupon) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO my_coupon(uid, coupon_id, status) VALUES (?, ?, ?)", m.UID, m.CouponID, m.Status).Exec()
}

func ModifyMyCoupon(m *MyCoupon) {
	o := orm.NewOrm()
	o.Raw("UPDATE my_coupon SET uid = ?, coupon_id = ?, status = ? WHERE id = ?", m.UID, m.CouponID, m.Status, m.ID).Exec()
}

func GetMyCouponByUIDAndCouponID(UID, couponID int) MyCoupon {
	var coupon MyCoupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM my_coupon WHERE uid = ? AND coupon_id = ?", UID, couponID).QueryRow(&coupon)
	return coupon
}

func GetMyCouponList() []MyCoupon {
	var coupons []MyCoupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM my_coupon").QueryRows(&coupons)
	return coupons
}

func GetMyCouponByCouponID(couponID int) MyCoupon {
	var m MyCoupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM my_coupon WHERE coupon_id = ?", couponID).QueryRow(&m)
	return m
}

func GetMyCouponByUID(UID int) []MyCoupon {
	var myCoupons []MyCoupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM my_coupon WHERE uid = ?", UID).QueryRows(&myCoupons)
	return myCoupons
}

func GetMyCouponByUIDAndStatus(UID, status int) []MyCoupon {
	var myCoupons []MyCoupon
	o := orm.NewOrm()
	o.Raw("SELECT * FROM my_coupon WHERE uid = ? AND status = 0", UID).QueryRows(&myCoupons)
	return myCoupons
}

func NumberOfMyCouponByUID(UID int) int {
	var count int
	o := orm.NewOrm()
	o.Raw("SELECT count(id) FROM my_coupon WHERE uid = ?", UID).QueryRow(&count)
	return count
}
