package controllers

import (
	"ChiMu/basic"
	"ChiMu/services"
)

type TeamController struct {
	basic.BasicController
}

func (c *TeamController) Team() {
	fid, _ := c.GetInt("fid")
	teams := services.GetTeamByFID(fid)
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: teams}
	c.ServeJSON()
}
