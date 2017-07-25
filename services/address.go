package services

import "ChiMu/models"

func AddAddress(a *models.Address) {
	models.AddAddress(a)
}

func DeleteAddress(ID int) {
	models.DeleteAddress(ID)
}

func GetAddressByUID(UID int) []models.Address {
	return models.GetAddressByUID(UID)
}

func GetAddressByID(ID int) models.Address {
	return models.GetAddressByID(ID)
}

func ModifyAddressByID(a *models.Address) {
	models.ModifyAddressByID(a)
}

func GetSelectedAddressByUID(UID int) models.Address {
	return models.GetSelectedAddressByUID(UID)
}
