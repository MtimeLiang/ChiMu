package services

import "ChiMu/models"

func GetProductWithID(ID int) models.Product {
	product := models.GetProductWithID(ID)
	product.Point = int(product.Price) / 10
	product.Images = models.GetImagesByPID(product.ID, 0)
	product.DescImages = models.GetImagesByPID(product.ID, 1)
	return product
}

func GetProductList() []models.Product {
	return models.GetProductList()
}

func AddProduct(p *models.Product) {
	models.AddProduct(p)
}
