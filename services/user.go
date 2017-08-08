package services

import "ChiMu/models"

func GetUserList() []models.User {
	return models.GetUserList()
}

func GetCommentUserByID(ID int) models.User {
	return models.GetCommentUserByID(ID)
}

func NumberOfCouponByUID(UID int) int {
	return models.NumberOfCouponByUID(UID)
}

func GetUserByID(ID int) models.User {
	return models.GetUserByID(ID)
}
