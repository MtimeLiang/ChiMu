package models

type Image struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	ProductID   int    `json:"product_id"`
	ProductType int    `json:"product_type"`
	BannerID    int    `json:"banner_id"`
}
