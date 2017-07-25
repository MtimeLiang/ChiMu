package services

import "ChiMu/models"

func GetCouponByPID(PID int) []models.Coupon {
	coupons := models.GetCouponByPID(PID)
	for c := range coupons {
		c.Title = "满" + c.MaxPrice + "减" + c.Price
	}
	return coupons
}

func GetCouponByUID(UID int) []models.Coupon {
	return models.GetCouponByUID(UID)
}
