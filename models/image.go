package models

import (
	"github.com/astaxie/beego/orm"
)

type Image struct {
	ID          int    `json:"id" orm:"column(id)"`
	URL         string `json:"url" orm:"column(url)"`
	ProductID   int    `json:"product_id" orm:"column(product_id)"`
	ProductType int    `json:"product_type"`
	BannerID    int    `json:"banner_id" orm:"column(banner_id)"`
	CommentID   int    `json:"comment_id" orm:"column(comment_id)"`
}

func AddImage(img Image) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO image(product_id, url, banner_id, product_type, comment_id) VALUES (?, ?, ?, ?, ?)",
		img.ProductID, img.URL, img.BannerID, img.ProductType, img.CommentID).Exec()
}

func GetImagesByPID(PID, productType int) []Image {
	var images []Image
	o := orm.NewOrm()
	o.Raw("SELECT id, product_id, product_type, url FROM image WHERE product_id = ? AND product_type = ?", PID, productType).QueryRows(&images)
	return images
}

func DeleteImageByPID(PID int) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM image WHERE product_id = ?", PID).Exec()
}

func GetImagesByBID(BID int) []Image {
	var images []Image
	o := orm.NewOrm()
	o.Raw("DELETE FROM image WHERE banner_id = ?", BID).QueryRows(&images)
	return images
}

func DeleteImageByBID(BID int) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM image WHERE banner_id = ?", BID).Exec()
}

func DeleteImageByID(ID int) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM image WHERE id = ?", ID).Exec()
}

func ModifyImageByID(img Image) {
	o := orm.NewOrm()
	o.Raw("UPDATE image SET url = ? WHERE id = ?", img.URL, img.ID).Exec()
}
