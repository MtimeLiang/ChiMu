package services

import "ChiMu/models"

func AddOrderDetail(detail *models.OrderDetail) {
	models.AddOrderDetail(detail)
}

func GetOrderDetailByOID(OID int) []models.OrderDetail {
	return models.GetOrderDetailByOID(OID)
}
