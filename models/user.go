package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	ID     int `json:"id"`
	Avatar string
	Name   string
	IsVip  int
	Point  int
	Award  int
	OpenID string `json:"openid"`
	Admin  string
}

func AddUser(u *User) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO user(avatar, name, isVip, point, award, openid, admin) VALUES(?, ?, ?, ?, ?, ?, ?)",
		u.Avatar, u.Name, u.IsVip, u.Point, u.Award, u.OpenID, u.Admin).Exec()
}

func ModifyUser(u *User) {
	o := orm.NewOrm()
	o.Raw("UPDATE user SET avatar = ?, name = ?, isVip = ?, point = ?, award = ?, openid = ?, admin = ? WHERE id = ?",
		u.Avatar, u.Name, u.IsVip, u.Point, u.Award, u.OpenID, u.Admin, u.ID).Exec()
}

func GetUserList() []User {
	var users []User
	o := orm.NewOrm()
	o.Raw("SELECT * FROM user").QueryRows(&users)
	return users
}

func GetUserWithOpenID(openID string) User {
	var user User
	o := orm.NewOrm()
	o.Raw("SELECT * FROM  user WHERE openid = ?", openID).QueryRow(&user)
	return user
}

func GetCommentUserByID(ID int) User {
	var user User
	o := orm.NewOrm()
	o.Raw("SELECT name, avatar FROM user WHERE id = ?", ID).QueryRow(&user)
	return user
}
