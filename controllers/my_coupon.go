package controllers

import (
	"ChiMu/basic"
	"ChiMu/models"
	"ChiMu/services"
	"ChiMu/utils"
)

type MyCouponAddController struct {
	basic.BasicController
}

type MyCouponController struct {
	basic.BasicController
}

func (c *MyCouponAddController) Get() {
	couponID, _ := c.GetInt("id")
	uid, _ := c.GetInt("uid")

	myCoupon := services.GetMyCouponByUIDAndCouponID(uid, couponID)
	if !utils.IsNil(myCoupon) {
		c.Data["json"] = basic.ResInfo{InfoMsg: "该用户已经领取", Status: 0, Data: nil}
		c.ServeJSON()
	}
	coupon := services.GetCouponByID(couponID)
	if utils.IsNil(coupon) {
		c.Data["json"] = basic.ResInfo{InfoMsg: "没有这张优惠券", Status: 0, Data: nil}
		c.ServeJSON()
	}
	myNewCoupon := new(models.MyCoupon)
	myNewCoupon.UID = uid
	myNewCoupon.CouponID = couponID
	myNewCoupon.Status = 0
	services.AddMyCoupon(myNewCoupon)
	c.Data["json"] = basic.ResInfo{InfoMsg: "添加成功", Status: 1, Data: nil}
	c.ServeJSON()
}

func (c *MyCouponController) Get() {
	uid, _ := c.GetInt("uid")
	status, statusErr := c.GetInt("status")

	if statusErr == nil {
		myCoupons := services.GetMyCouponByUIDAndStatus(uid, status)
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: myCoupons}
		c.ServeJSON()
	}
	myCoupons := services.GetMyCouponByUID(uid)
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: myCoupons}
	c.ServeJSON()
}
