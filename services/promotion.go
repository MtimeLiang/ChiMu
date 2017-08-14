package services

import "ChiMu/models"

func GetPromotionList(PID int) []models.Promotion {
	return models.GetPromotionList(PID)
}

func GetPromotionWithID(ID int) models.Promotion {
	return models.GetPromotionWithID(ID)
}

func GetPromotions() []models.Promotion {
	return models.GetPromotions()
}

func ModifyPromotionWithID(p *models.Promotion) {
	models.ModifyPromotionWithID(p)
}

func DeletePromotionWithID(ID int) {
	models.DeletePromotionWithID(ID)
}
