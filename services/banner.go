package services

import (
	"ChiMu/models"
	"ChiMu/utils"
	"mime/multipart"
)

func AddBanner(b *models.Banner) {
	models.AddBanner(b)
}

func ModifyBanner(b *models.Banner, f multipart.File, h multipart.FileHeader) {
	utils.RemoveFile(b.ImgURL)
	imgURL, _ := utils.SaveFile(f, h)
	if len(imgURL) > 0 {
		b.ImgURL = imgURL
	}
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
