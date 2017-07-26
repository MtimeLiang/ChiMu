package models

import (
	"github.com/astaxie/beego/orm"
)

type Banner struct {
	ID     int    `orm:"column(id)"`
	ImgURL string `json:"imgurl" orm:"column(imgurl)"`
	URL    string `json:"url" orm:"column(url)"`
	Title  string
}

func GetBannerList() []Banner {
	var banners []Banner
	o := orm.NewOrm()
	o.Raw("SELECT * FROM banner").QueryRows(banners)
	return banners
}

func GetBannerByID(ID int) Banner {
	var banner Banner
	o := orm.NewOrm()
	o.Raw("SELECT * FROM banner WHERE id = ?", ID).QueryRow(banner)
	return banner
}

func AddBanner(b *Banner) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO banner(title, url, imgurl) VALUES (?, ?, ?)", b.Title, b.URL, b.ImgURL).Exec()
}

func ModifyBanner(b *Banner) {
	o := orm.NewOrm()
	o.Raw("UPDATE banner SET title = ?, url = ?, imgurl = ? WHERE id = ?", b.Title, b.URL, b.ImgURL, b.ID).Exec()
}

func DeleteBanner(ID int) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM banner WHERE id = ?", ID).Exec()
}
