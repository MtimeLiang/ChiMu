package controllers

import (
	"ChiMu/basic"
	"ChiMu/services"
)

type PointController struct {
	basic.BasicController
}

func (c *PointController) Get() {
	uid, _ := c.GetInt("uid")

	points := services.GetPointListByUID(uid)
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: points}
	c.ServeJSON()
}
