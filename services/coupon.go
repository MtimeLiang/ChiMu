package services

import "ChiMu/models"
import "strconv"

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

func GetCouponByPIDAndUID(PID int, UID int) []models.Coupon {
	return models.GetCouponByPIDAndUID(PID, UID)
}
