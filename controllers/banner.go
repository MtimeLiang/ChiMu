package controllers

import (
	"ChiMu/basic"
	"ChiMu/services"
)

type BannerListController struct {
	basic.BasicController
}

type BannerDetailController struct {
	basic.BasicController
}

type BannerAddController struct {
	basic.BasicController
}

func (c *BannerListController) Get() {
	banners := services.GetBannerList()
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取banner成功", Status: 1, Data: banners}
	c.ServeJSON()
}

func (c *BannerDetailController) Get() {
	id, _ := c.GetInt("id")
	banner := services.GetBannerByID(id)
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取banner成功", Status: 1, Data: banner}
	c.ServeJSON()
}

func (c *BannerAddController) Post() {
	c.SaveFile("file")
	c.Data["json"] = basic.ResInfo{InfoMsg: "添加banner成功", Status: 1, Data: nil}
	c.ServeJSON()
}
