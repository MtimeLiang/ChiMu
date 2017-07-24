package models

type FeeRecord struct {
	ID       int    `json:"id"`
	UID      string `json:"uid"`
	Money    float64
	FeeMoney float64
}
