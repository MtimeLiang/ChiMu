package services

import (
	"ChiMu/models"
	"time"
)

func AddOrder(o *models.Order) {
	models.AddOrder(o)
}

func GetOrderList() []models.Order {
	orders := models.GetOrderList()
	for _, o := range orders {
		configOrderDetail(&o)
	}
	return orders
}

func GetOrderByID(ID int) models.Order {
	o := models.GetOrderByID(ID)
	configOrderDetail(&o)
	return o
}

func GetOrderByUID(UID int) []models.Order {
	orders := models.GetOrderByUID(UID)
	for _, o := range orders {
		configOrderDetail(&o)
	}
	return orders
}

func GetOrderByOrderNum(orderNum string) models.Order {
	o := models.GetOrderByOrderNum(orderNum)
	// return configOrderDetail(o)
	configOrderDetail(&o)
	return o
}

func ModifyOrder(o *models.Order) {
	if o.Status == 1 {
		// 订单完成
		p := new(models.Point)
		p.Point = 10
		p.Description = "买酒"
		p.CreateTime = time.Now()
		p.UID = o.UID
		models.AddPoint(p)
	}
	models.ModifyOrder(o)
}

func GetOrderByUIDAndStatus(UID, status int) []models.Order {
	return models.GetOrderByUIDAndStatus(UID, status)
}

func configOrderDetail(o *models.Order) {
	// 地址信息
	if o.AddressID == 0 {
		o.AddressInfo = GetSelectedAddressByUID(o.UID)
	} else {
		o.AddressInfo = GetAddressByID(o.AddressID)
	}
	// 优惠券信息
	if o.CouponID > 0 {
		o.CouponInfo = models.GetCouponByID(o.CouponID)
	}
	// 订单详情信息
	o.OrderDetails = GetOrderDetailByOID(o.ID)
	for _, detail := range o.OrderDetails {
		detail.ProductInfo = models.GetProductWithID(detail.PID)
	}
}
