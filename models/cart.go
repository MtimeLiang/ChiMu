package models

type Cart struct {
	ID         int `json:"id"`
	UID        int `json:"uid"`
	PID        int `json:"pid"`
	Count      int
	Title      string
	SubMessage string `json:"submessage"`
	Price      float64
	Volume     int
	Images     string
}
