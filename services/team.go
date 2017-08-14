package services

import (
	"ChiMu/models"
)

func GetTeamByFID(FID int) []models.Team {
	return models.GetTeamByFID(FID)
}
