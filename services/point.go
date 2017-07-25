package services

import "ChiMu/models"

func GetPointListByUID(UID int) []models.Point {
	return models.GetPointListByUID(UID)
}

func AddPoint(p *models.Point) {
	models.AddPoint(p)
}
