package models

import (
	"github.com/astaxie/beego/orm"
)

type Product struct {
	ID               int         `json:"id" orm:"column(id)"`
	Title            string      `json:"title"`
	SubMessage       string      `json:"submessage" orm:"column(submessage)"`
	Price            float64     `json:"price"`
	Volume           int         `json:"colume"`
	Image            string      `json:"image"`
	Description      string      `json:"description"`
	OriginPrice      float64     `json:"origprice" orm:"column(origprice)"`
	Count            int         `json:"count"`
	DescriptionImage string      `json:"description_image"`
	FreightMoney     float64     `json:"freight_money"`
	Point            int         `json:"point"`
	Sales            int         `json:"sales"`
	Images           []Image     `json:"images"`
	DescImages       []Image     `json:"desc_images"`
	Invalid          int         `json:"invalid"`    // 默认下架
	Promotions       []Promotion `json:"promotions"` // 促销列表
	Coupons          []Coupon    `json:"coupons"`    // 优惠券列表
	Comments         []Comment   `json:"comments"`   // 评论列表
}

func AddProduct(p *Product) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO product(title, submessage, price, volume, image, description, origprice, count, description_image, freight_money, point, sales, invalid) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?",
		p.Title, p.SubMessage, p.Price, p.Volume, p.Image, p.Description, p.OriginPrice, p.Count, p.DescriptionImage, p.FreightMoney, p.Point, p.Sales, p.Invalid).Exec()
}

func GetProductList() []Product {
	var products []Product
	o := orm.NewOrm()
	o.Raw("SELECT * FROM product").QueryRows(&products)
	return products
}

func GetProductWithID(ID int) Product {
	var product Product
	o := orm.NewOrm()
	o.Raw("SELECT * FROM product WHERE id = ?", ID).QueryRow(&product)
	return product
}

func ModifyProduct(p *Product) {
	o := orm.NewOrm()
	o.Raw("UPDATE product SET title = ?, submessage = ?, price = ?, volume = ?, description = ?, origprice = ?, count = ?, freight_money = ?, point = ?, sales = ?, image = ?, description_image = ?, invalid = ? WHERE id = ?",
		p.Title, p.SubMessage, p.Price, p.Volume, p.Description, p.OriginPrice, p.Count, p.FreightMoney, p.Point, p.Sales, p.Image, p.DescriptionImage, p.Invalid, p.ID).Exec()
}
