package services

import (
	"ChiMu/models"
	"ChiMu/utils"
	"strconv"
)

func GetCouponList() []models.Coupon {
	coupons := models.GetCouponList()
	for _, c := range coupons {
		c.Title = "满" + strconv.Itoa(c.MaxPrice) + "减" + strconv.Itoa(c.Price)
	}
	return coupons
}

func GetCouponByID(ID int) models.Coupon {
	return models.GetCouponByID(ID)
}

func GetCouponByPID(PID int) []models.Coupon {
	coupons := models.GetCouponByPID(PID)
	for _, c := range coupons {
		c.Title = "满" + strconv.Itoa(c.MaxPrice) + "减" + strconv.Itoa(c.Price)
	}
	return coupons
}

func GetCouponByUID(UID int) []models.Coupon {
	return models.GetCouponByUID(UID)
}

func AddCoupon(c *models.Coupon) {
	models.AddCoupon(c)
}

func ModifyCouponByID(c *models.Coupon) {
	models.ModifyCouponByID(c)
}

func DeleteCouponByID(ID int) {
	coupon := models.GetCouponByID(ID)
	coupon.IsDelete = 1
	models.ModifyCouponByID(&coupon)
}

func GetCouponByPIDAndUID(pid, uid int) []models.Coupon {
	coupons := models.GetCouponByPID(pid)
	for _, c := range coupons {
		c.Title = "满" + strconv.Itoa(c.MaxPrice) + "减" + strconv.Itoa(c.Price)
		// 获取用户的使用情况
		m := models.GetMyCouponByUIDAndCouponID(uid, c.ID)
		if utils.IsNil(m) {
			c.Status = 1
		} else {
			c.Status = 0
		}
	}
	return coupons
}
