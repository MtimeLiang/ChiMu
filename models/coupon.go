package models

type Coupon struct {
	ID        int `json:"id"`
	PID       int `json:"pid"`
	UID       int `json:"uid"`
	Price     int
	MaxPrice  int `json:"max_price"`
	Title     string
	BuildTime string `json:"build_time"`
	EndTime   string `json:"end_time"`
}
