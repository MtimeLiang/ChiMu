package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Point struct {
	ID          int `json:"id"`
	UID         int `json:"uid"`
	Point       int
	CreateTime  time.Time `json:"create_time"`
	Type        int
	Description string
}

func GetPointListByUID(UID int) []Point {
	var points []Point
	o := orm.NewOrm()
	o.Raw("SELECT * FROM point WHERE uid = ? ORDER BY create_time ASC", UID).QueryRows(&points)
	return points
}

func AddPoint(p *Point) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO point(uid, point, create_time, type, desciption) VALUES (?, ?, ?, ?, ?)",
		p.UID, p.Point, p.CreateTime, p.Type, p.Description).Exec()
}
