package models

import (
	"time"
)

type Cash struct {
	ID    int `json:"id":`
	UID   int `json:"uid"`
	Money float64
	Date  time.Time
}
