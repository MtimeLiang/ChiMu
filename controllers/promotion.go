package controllers

import (
	"ChiMu/basic"
	"ChiMu/services"
)

type PromotionController struct {
	basic.BasicController
}

func (c *PromotionController) Promotion() {
	id, idErr := c.GetInt("id")
	pid, pidErr := c.GetInt("pid")

	if idErr == nil {
		promotion := services.GetPromotionWithID(id)
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: promotion}
		c.ServeJSON()
	}
	if pidErr == nil {
		promotions := services.GetPromotionList(pid)
		c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: promotions}
		c.ServeJSON()
	}
	promotions := services.GetPromotions()
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: promotions}
	c.ServeJSON()
}

func (c *PromotionController) ModifyPromotion() {
	id, _ := c.GetInt("id")
	pid, pidErr := c.GetInt("pid")
	name := c.GetString("name")
	desc := c.GetString("description")

	promotion := services.GetPromotionWithID(id)
	if pidErr == nil {
		promotion.PID = pid
	}
	if len(name) > 0 {
		promotion.Name = name
	}
	if len(desc) > 0 {
		promotion.Description = desc
	}
	services.ModifyPromotionWithID(&promotion)
	c.Data["json"] = basic.ResInfo{InfoMsg: "修改成功", Status: 1, Data: promotion}
	c.ServeJSON()
}

func (c *PromotionController) DeletePromotion() {
	id, _ := c.GetInt("id")
	services.DeletePromotionWithID(id)
	c.Data["json"] = basic.ResInfo{InfoMsg: "删除成功", Status: 1, Data: nil}
	c.ServeJSON()
}
