package controllers

import (
	"ChiMu/basic"
	"ChiMu/models"
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

type BannerModifyController struct {
	basic.BasicController
}

type BannerDeleteController struct {
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
	title := c.GetString("title")
	url := c.GetString("url")

	if imgURL, _ := c.SaveFile("file"); len(imgURL) > 0 {
		banner := new(models.Banner)
		banner.Title = title
		banner.URL = url
		banner.ImgURL = imgURL
		services.AddBanner(banner)
		c.Data["json"] = basic.ResInfo{InfoMsg: "添加banner成功", Status: 1, Data: nil}
		c.ServeJSON()
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "添加banner失败", Status: 0, Data: nil}
	c.ServeJSON()
}

func (c *BannerModifyController) Post() {
	id, _ := c.GetInt("id")
	title := c.GetString("title")
	url := c.GetString("url")

	if imgURL, _ := c.SaveFile("file"); len(imgURL) > 0 {
		banner := new(models.Banner)
		banner.ID = id
		banner.Title = title
		banner.URL = url
		banner.ImgURL = imgURL
		services.ModifyBanner(banner)
		c.Data["json"] = basic.ResInfo{InfoMsg: "修改banner成功", Status: 1, Data: nil}
		c.ServeJSON()
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "修改banner失败", Status: 0, Data: nil}
	c.ServeJSON()
}

func (c *BannerDeleteController) Get() {
	id, _ := c.GetInt("id")
	services.DeleteBanner(id)
	c.Data["json"] = basic.ResInfo{InfoMsg: "删除banner成功", Status: 1, Data: nil}
	c.ServeJSON()
}
