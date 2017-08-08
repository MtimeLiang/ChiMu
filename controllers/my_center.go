package controllers

import (
	"ChiMu/basic"
	"ChiMu/services"
)

type MyUserInfoController struct {
	basic.BasicController
}

type MyOrderController struct {
	basic.BasicController
}

func (c *MyUserInfoController) Get() {
	id, _ := c.GetInt("id")

	user := services.GetUserByID(id)
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: user}
	c.ServeJSON()
}

func (c *MyOrderController) Get() {
	uid, _ := c.GetInt("uid")
	status, statusErr := c.GetInt("status")

	if statusErr == nil {
		orders := services.GetOrderByUIDAndStatus(uid, status)
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: orders}
		c.ServeJSON()
	}
	orders := services.GetOrderByUID(uid)
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: orders}
	c.ServeJSON()
}
