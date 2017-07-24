package models

import (
	"time"
)

type Point struct {
	ID          int `json:"id"`
	UID         int `json:"uid"`
	Point       int
	CreateTime  time.Time `json:"create_time"`
	Type        int
	Description string
}
