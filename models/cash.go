package models

import (
	"time"
)

type Cash struct {
	ID    int `json:"id" orm:"column(id)"`
	UID   int `json:"uid" orm:"column(uid)"`
	Money float64
	Date  time.Time
}
