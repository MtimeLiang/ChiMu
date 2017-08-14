package models

import (
	"github.com/astaxie/beego/orm"
)

type Promotion struct {
	ID          int    `json:"id" orm:"column(id)"`
	PID         int    `json:"pid" orm:"column(pid)"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetPromotionList(PID int) []Promotion {
	var promotions []Promotion
	o := orm.NewOrm()
	o.Raw("SELECT * FROM promotion WHERE pid = ?", PID).QueryRows(&promotions)
	return promotions
}

func GetPromotionWithID(ID int) Promotion {
	var promotion Promotion
	o := orm.NewOrm()
	o.Raw("SELECT * FROM promotion WHERE id = ?", ID).QueryRow(&promotion)
	return promotion
}

func GetPromotions() []Promotion {
	var promotions []Promotion
	o := orm.NewOrm()
	o.Raw("SELECT * FROM promotion").QueryRow(&promotions)
	return promotions
}

func ModifyPromotionWithID(p *Promotion) {
	o := orm.NewOrm()
	o.Raw("UPDATE promotion SET pid = ?, name = ?, description = ? WHERE id = ?", p.PID, p.Name, p.Description, p.ID).Exec()
}

func DeletePromotionWithID(ID int) {
	o := orm.NewOrm()
	o.Raw("DELETE FROM promotion WHERE id = ?", ID).Exec()
}
