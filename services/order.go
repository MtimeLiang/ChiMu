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
	for o := range orders {
		o = configOrderDetail(o)
	}
	return orders
}

func GetOrderByID(ID int) models.Order {
	o := models.GetOrderByID(ID)
	return configOrderDetail(o)
}

func GetOrderByUID(UID int) []models.Order {
	orders := models.GetOrderByUID(UID)
	for o := range orders {
		o = configOrderDetail(o)
	}
	return orders
}

func GetOrderByOrderNum(orderNum string) models.Order {
	o := models.GetOrderByOrderNum(orderNum)
	return configOrderDetail(o)
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

func configOrderDetail(o *models.Order) models.Order {
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
	for detail := range o.OrderDetails {
		detail.ProductInfo = models.GetProductWithID(o.PID)
	}
	return o
}
