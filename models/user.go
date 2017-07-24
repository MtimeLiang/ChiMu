package models

type User struct {
	ID     int `json:"id"`
	Avatar string
	Name   string
	IsVip  int
	Point  int
	Award  int
	OpenID string `json:"openid"`
}
