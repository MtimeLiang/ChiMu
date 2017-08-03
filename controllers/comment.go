package controllers

import (
	"ChiMu/basic"
	"ChiMu/models"
	"ChiMu/services"
	"time"
)

type CommentController struct {
	basic.BasicController
}

type CommentAddController struct {
	basic.BasicController
}

func (c *CommentController) Get() {
	pid, _ := c.GetInt("pid")
	comments := services.GetCommentByPID(pid)
	for _, comment := range comments {
		user := services.GetCommentUserByID(comment.UID)
		comment.UserName = user.Name
		comment.UserAvatar = user.Avatar
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取评论列表成功", Status: 1, Data: comments}
	c.ServeJSON()
}

func (c *CommentAddController) Post() {
	title := c.GetString("title")
	pid, _ := c.GetInt("pid")
	uid, _ := c.GetInt("uid")

	comment := new(models.Comment)
	comment.Title = title
	comment.UID = uid
	comment.PID = pid
	comment.CreateTime = time.Now()

	services.AddComment(comment)

	if imgURLs, err := c.SaveFiles("file"); err == nil {
		for _, imgURL := range imgURLs {
			image := new(models.Image)
			image.URL = imgURL
			image.ProductID = 0
			image.ProductType = 0
			image.BannerID = 0
			image.CommentID = comment.ID
			models.AddImage(*image)
		}
		c.Data["json"] = basic.ResInfo{InfoMsg: "评论成功", Status: 1, Data: nil}
		c.ServeJSON()
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "评论失败", Status: 0, Data: nil}
	c.ServeJSON()
}
