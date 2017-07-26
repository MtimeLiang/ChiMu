package models

type FeeRecord struct {
	ID       int     `json:"id" orm:"column(id)"`
	UID      string  `json:"uid" orm:"column(uid)"`
	Money    float64 `json:"money"`
	FeeMoney float64 `json:"feeMoney"`
}
