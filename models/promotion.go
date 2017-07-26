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
