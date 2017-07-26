package models

type Team struct {
	ID    int    `json:"id" orm:"column(id)"`
	UID   string `json:"uid" orm:"column(uid)"`
	FriID string `json:"friId" orm:"column(fri_id)"`
}
