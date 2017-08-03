package controllers

import (
	"ChiMu/basic"
	"ChiMu/models"
	"ChiMu/services"
	"time"
)

type CouponController struct {
	basic.BasicController
}

type CouponAddController struct {
	basic.BasicController
}

type CouponModifyController struct {
	basic.BasicController
}

type CouponDeleteController struct {
	basic.BasicController
}

func (c *CouponController) Get() {
	id, idErr := c.GetInt("id")
	pid, pidErr := c.GetInt("pid")
	uid, uidErr := c.GetInt("uid")

	if idErr == nil {
		coupon := services.GetCouponByID(id)
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: coupon}
	} else if pidErr == nil && uidErr == nil {
		coupons := services.GetCouponByPIDAndUID(pid, uid)
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: coupons}
	} else {
		coupons := services.GetCouponList()
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: coupons}
	}
	c.ServeJSON()
}

func (c *CouponAddController) Get() {
	coupon := getCouponWithData(&c.BasicController, true)
	services.AddCoupon(&coupon)
	c.Data["json"] = basic.ResInfo{InfoMsg: "添加优惠券成功", Status: 1, Data: nil}
	c.ServeJSON()
}

func (c *CouponModifyController) Get() {
	coupon := getCouponWithData(&c.BasicController, false)
	services.ModifyCouponByID(&coupon)
	c.Data["json"] = basic.ResInfo{InfoMsg: "修改优惠券成功", Status: 1, Data: nil}
	c.ServeJSON()
}

func (c *CouponDeleteController) Get() {
	id, _ := c.GetInt("id")
	services.DeleteCouponByID(id)
	c.Data["json"] = basic.ResInfo{InfoMsg: "删除优惠券成功", Status: 1, Data: nil}
	c.ServeJSON()
}

// Private Method
func getCouponWithData(c *basic.BasicController, isAdd bool) models.Coupon {
	price, priceErr := c.GetInt("price")
	maxPrice, maxPriceErr := c.GetInt("max_price")
	pid, pidErr := c.GetInt("pid")
	id, _ := c.GetInt("id")
	day, _ := c.GetInt64("day")

	var coupon models.Coupon
	if isAdd {
		coupon = *new(models.Coupon)
		coupon.IsDelete = 0
		t := time.Now()
		coupon.BuildTime = t.Format("2016-09-01")
		d, _ := time.ParseDuration("+24h")
		coupon.EndTime = t.Add(d * time.Duration(day)).Format("2016-09-01")
	} else {
		coupon = services.GetCouponByID(id)
	}

	if priceErr == nil {
		coupon.Price = price
	} else if maxPriceErr == nil {
		coupon.MaxPrice = maxPrice
	} else if pidErr == nil {
		coupon.PID = pid
	}
	return coupon
}
