package models

import (
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	ID         int `json:"id"`
	PID        int `json:"pid"`
	UID        int `json:"uid"`
	Date       string
	Title      string
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
}

func GetCommentByPID(PID int) []Comment {
	var comments []Comment
	o := orm.NewOrm()
	o.Raw("SELECT * FROM comment WHERE pid = ?", PID).QueryRow(&comments)
	return comments
}
