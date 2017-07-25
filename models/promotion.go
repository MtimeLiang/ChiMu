package models

import (
	"github.com/astaxie/beego/orm"
)

type Promotion struct {
	ID          int `json:"id"`
	PID         int `json:"pid"`
	Name        string
	Description string
}

func GetPromotionList(PID int) []Promotion {
	var promotions []Promotion
	o := orm.NewOrm()
	o.Raw("SELECT * FROM promotion WHERE pid = ?", PID).QueryRows(&promotions)
	return promotions
}
