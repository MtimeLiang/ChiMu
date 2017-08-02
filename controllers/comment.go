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

func (c *CommentAddController) Get() {
	title := c.GetString("title")
	pid, _ := c.GetInt("pid")
	uid, _ := c.GetInt("uid")

	comment := new(models.Comment)
	comment.Title = title
	comment.UID = uid
	comment.PID = pid
	comment.CreateTime = time.Now()
	/// Add Save Images Methods.
}

// GetFiles return multi-upload files
// files, err:=c.GetFiles("myfiles")
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusNoContent)
//		return
//	}
// for i, _ := range files {
//	//for each fileheader, get a handle to the actual file
//	file, err := files[i].Open()
//	defer file.Close()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	//create destination file making sure the path is writeable.
//	dst, err := os.Create("upload/" + files[i].Filename)
//	defer dst.Close()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	//copy the uploaded file to the destination file
//	if _, err := io.Copy(dst, file); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
// }
