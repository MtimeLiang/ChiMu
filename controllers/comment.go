package controllers

import (
	"ChiMu/basic"
	"ChiMu/services"
)

type CommentController struct {
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
