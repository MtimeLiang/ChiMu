package models

import "github.com/astaxie/beego/orm"

type Team struct {
	ID    int    `json:"id" orm:"column(id)"`
	UID   string `json:"uid" orm:"column(uid)"`
	FriID string `json:"friId" orm:"column(fri_id)"`
}

func GetTeamByFID(FID int) []Team {
	var teams []Team
	o := orm.NewOrm()
	o.Raw("SELECT * FROM team WHERE fri_id = ?", FID).QueryRows(&teams)
	return teams
}
