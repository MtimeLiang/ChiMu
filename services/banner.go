package services

import (
	"ChiMu/models"
)

func AddBanner(b *models.Banner) {
	models.AddBanner(b)
}

func ModifyBanner(b *models.Banner) {
	models.ModifyBanner(b)
}

func GetBannerList() []models.Banner {
	return models.GetBannerList()
}

func GetBannerByID(ID int) models.Banner {
	return models.GetBannerByID(ID)
}

func DeleteBanner(ID int) {
	models.DeleteBanner(ID)
}
