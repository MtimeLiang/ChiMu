package controllers

import (
	"ChiMu/basic"
	"ChiMu/models"
	"ChiMu/services"
)

type AddressAddController struct {
	basic.BasicController
}

type AddressModifyController struct {
	basic.BasicController
}

func (c *AddressAddController) Get() {
	uid, err := c.GetInt("uid")
	name := c.GetString("name")
	tel := c.GetString("tel")
	address := c.GetString("address")
	provinceCity := c.GetString("province_city")

	if err != nil {
		c.Data["json"] = basic.ResInfo{InfoMsg: "参数错误", Status: 0, Data: ""}
		c.ServeJSON()
	}
	addressList := services.GetAddressByUID(uid)

	addressModel := new(models.Address)
	addressModel.UID = uid
	addressModel.Name = name
	addressModel.Tel = tel
	addressModel.Address = address
	addressModel.ProvinceCity = provinceCity

	if len(addressList) == 0 {
		// 第一次添加就是默认地址，默认选中
		addressModel.IsDefault = 1
		addressModel.IsSelected = 1
	} else {
		addressModel.IsDefault = 0
		addressModel.IsSelected = 0
	}
	services.AddAddress(addressModel)
	c.Data["json"] = basic.ResInfo{InfoMsg: "添加地址成功", Status: 1, Data: addressModel}
	c.ServeJSON()
}

func (c *AddressModifyController) Get() {
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	tel := c.GetString("tel")
	address := c.GetString("address")
	provinceCity := c.GetString("province_city")
	uid, _ := c.GetInt("uid")
	isSelected, _ := c.GetInt("selected")
	isDefault, _ := c.GetInt("is_default")

	addressModel := services.GetAddressByID(id)
	addressModel.Name = name
	addressModel.Tel = tel
	addressModel.Address = address
	addressModel.ProvinceCity = provinceCity

	if uid > 0 {
		addressList := services.GetAddressByUID(uid)
		if isSelected > 0 {
			for _, v := range addressList {
				v.IsSelected = 0
				services.ModifyAddressByID(&v)
			}
			addressModel.IsSelected = 1
		} else if isDefault > 0 {
			for _, v := range addressList {
				v.IsDefault = 0
				services.ModifyAddressByID(&v)
			}
			addressModel.IsDefault = 1
		}
	}
	services.ModifyAddressByID(&addressModel)
	c.Data["json"] = basic.ResInfo{InfoMsg: "修改成功", Status: 1, Data: addressModel}
	c.ServeJSON()
}
