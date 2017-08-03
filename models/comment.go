package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	ID         int       `json:"id" orm:"column(id)"`
	PID        int       `json:"pid" orm:"column(pid)"`
	UID        int       `json:"uid" orm:"column(uid)"`
	CreateTime time.Time `json:"create_time"`
	Title      string    `json:"title"`
	UserName   string    `json:"user_name"`
	UserAvatar string    `json:"user_avatar"`
}

func GetCommentByPID(PID int) []Comment {
	var comments []Comment
	o := orm.NewOrm()
	o.Raw("SELECT * FROM comment WHERE pid = ?", PID).QueryRow(&comments)
	return comments
}

func AddComment(c *Comment) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO comment(pid, title, uid, create_time) VALUES (?, ?, ?, ?)", c.PID, c.Title, c.UID, c.CreateTime).Exec()
}
