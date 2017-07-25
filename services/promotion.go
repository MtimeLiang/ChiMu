package services

import "ChiMu/models"

func GetPromotionList(PID int) []models.Promotion {
	return models.GetPromotionList(PID)
}
