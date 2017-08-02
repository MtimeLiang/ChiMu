package controllers

import (
	"ChiMu/basic"
	"ChiMu/services"
	"ChiMu/utils"
)

type CartAddController struct {
	basic.BasicController
}

type CartController struct {
	basic.BasicController
}

type CartDeleteController struct {
	basic.BasicController
}

func (c *CartAddController) Get() {
	pid, _ := c.GetInt("pid")
	uid, _ := c.GetInt("uid")
	count, _ := c.GetInt("count")

	cart := services.GetCart(pid, uid)
	if utils.IsNil(cart) {
		cart.Count = count + cart.Count
	} else {
		cart.PID = pid
		cart.UID = uid
		cart.Count = count
		services.AddCart(&cart)
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "添加成功", Status: 1, Data: cart}
	c.ServeJSON()
}

func (c *CartController) Get() {
	uid, _ := c.GetInt("uid")

	carts := services.GetCartByUID(uid)
	for _, cart := range carts {
		cart.SetProductModel(services.GetProductWithID(cart.PID))
	}
	c.Data["json"] = basic.ResInfo{InfoMsg: "获取成功", Status: 1, Data: carts}
	c.ServeJSON()
}

func (c *CartDeleteController) Get() {
	id, _ := c.GetInt("id")
	services.DeleteCart(id)
	c.Data["json"] = basic.ResInfo{InfoMsg: "删除成功", Status: 1, Data: nil}
	c.ServeJSON()
}
