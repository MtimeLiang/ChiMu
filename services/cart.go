package services

import "ChiMu/models"

func AddCart(c *models.Cart) {
	return models.AddCart(c)
}

func GetCart(PID, UID int) models.Cart {
	return models.GetCart(PID, UID)
}

func ModifyCartByID(c *models.Cart) {
	models.ModifyCartByID(c)
}

func GetCartByUID(UID int) []models.Cart {
	return models.GetCartByUID(UID)
}

func DeleteCart(ID int) {
	models.DeleteCart(ID)
}
