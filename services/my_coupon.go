package services

import "ChiMu/models"

func AddMyCoupon(m *models.MyCoupon) {
	models.AddMyCoupon(m)
}

func ModifyMyCoupon(m *models.MyCoupon) {
	models.ModifyMyCoupon(m)
}

func GetMyCouponList() []models.MyCoupon {
	myCoupons := models.GetMyCouponList()
	for _, m := range myCoupons {
		m.CouponInfo = models.GetCouponByID(m.ID)
	}
	return myCoupons
}

func GetMyCouponByUID(UID int) []models.MyCoupon {
	myCoupons := models.GetMyCouponByUID(UID)
	for _, m := range myCoupons {
		m.CouponInfo = models.GetCouponByID(m.ID)
	}
	return myCoupons
}

func GetMyCouponByUIDAndCouponID(UID, couponID int) models.MyCoupon {
	return models.GetMyCouponByUIDAndCouponID(UID, couponID)
}

func GetMyCouponByUIDAndStatus(UID, status int) []models.MyCoupon {
	myCoupons := models.GetMyCouponByUIDAndStatus(UID, status)
	for _, m := range myCoupons {
		m.CouponInfo = models.GetCouponByID(m.ID)
	}
	return myCoupons
}
