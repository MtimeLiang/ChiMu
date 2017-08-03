package services

import "ChiMu/models"

func GetCommentByPID(PID int) []models.Comment {
	return models.GetCommentByPID(PID)
}

func AddComment(c *models.Comment) {
	models.AddComment(c)
}
