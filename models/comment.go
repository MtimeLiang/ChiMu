package models

type Comment struct {
	ID         int `json:"id"`
	PID        int `json:"pid"`
	UID        int `json:"uid"`
	Date       string
	Title      string
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
}
