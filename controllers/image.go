package controllers

import (
	"ChiMu/basic"
)

type ImageController struct {
	basic.BasicController
}

func (c *ImageController) Get() {
	imgName := c.GetString("name")
	path := c.FileDir() + imgName
	c.Ctx.Output.Download(path)
}
