package services

import (
	"ChiMu/models"
	"mime/multipart"
)

func AddProduct(p *models.Product, fs []multipart.File, hs []multipart.FileHeader, fs1 []multipart.File, hs1 []multipart.FileHeader) {
	/// 稍后添加
}

func getImageList(images []models.Image, fs []multipart.File, hs []multipart.FileHeader, ID int, productType int) []models.Image {
	var image models.Image
	// 有删除图片
	if len(images) > len(fs) {
		for i := len(images) - 1; i > len(fs); i-- {
			tmp := images[i]
			models.DeleteImageByID(tmp.ID)
			remove(images, tmp)
		}
	}
	for i := 0; i < len(fs); i++ {
		f := fs[i]
		if f != nil {
			if len(images) > i {
				// remove
				URL := images[i]
			}
		}
	}
}
