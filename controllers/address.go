package controllers

import (
	"Akaki/models"
	"ChiMu/basic"
)

type AddressController struct {
	basic.BasicController
}

func (c *AddressController) Get() {
	uid, _ := c.GetInt("uid")
	name := c.GetString("name")
	tel := c.GetString("tel")
	address := c.GetString("address")
	provinceCity := c.GetString("provinceCity")

	// addressList := services.GetAddressByUID(uid)

	addressModel := new(models.Address)
	addressModel.UID = uid
	addressModel.Name = name
	addressModel.Tel = tel
	addressModel.Address = address
	addressModel.ProvinceCity = provinceCity
}
