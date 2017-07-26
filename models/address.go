package models

import (
	"github.com/astaxie/beego/orm"
)

type Address struct {
	ID           int
	UID          int
	Name         string
	Tel          string
	Address      string
	ProvinceCity string `json:"province_city"`
	IsDefault    int    `json:"is_default"`
	IsSelected   int    `json:"is_selected"`
	IsDelete     int    `json:"is_delete"`
}

func AddAddress(a *Address) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO address(name, tel, uid, address, province_city, is_default, is_selected) VALUES(?, ?, ?, ?, ?, ?, ?)",
		a.Name, a.Tel, a.UID, a.Address, a.ProvinceCity, a.IsDefault, a.IsSelected).Exec()
}

func DeleteAddress(ID int) {
	o := orm.NewOrm()
	o.Raw("UPDATE address SET is_delete = 1 WHERE id = ?", ID)
}

func GetAddressByUID(UID int) []Address {
	var address []Address
	o := orm.NewOrm()
	o.Raw("SELECT * FROM address WHERE uid = ? AND is_delete = 0", UID).QueryRow(&address)
	return address
}

func GetAddressByID(ID int) Address {
	var address Address
	o := orm.NewOrm()
	o.Raw("SELECT * FROM address WHERE id = ?", ID).QueryRow(&address)
	return address
}

func ModifyAddressByID(a *Address) {
	o := orm.NewOrm()
	o.Raw("UPDATE address SET name = ?, tel = ?, address = ?, province_city = ?, is_default = ?, is_selected = ? WHERE id = ?",
		a.Name, a.Tel, a.Address, a.ProvinceCity, a.IsDefault, a.IsSelected, a.ID).Exec()
}

func GetSelectedAddressByUID(UID int) Address {
	var address Address
	o := orm.NewOrm()
	o.Raw("SELECT * FROM address WHERE is_selected = 1 AND uid = ?", UID).QueryRow(&address)
	return address
}
