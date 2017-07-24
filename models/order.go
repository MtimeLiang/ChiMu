package models

import (
	"time"
)

type Order struct {
	ID         int    `json:"id"`
	UID        int    `json:"uid"`
	OrderNum   string `json:"order_num"`
	Status     int
	Count      int
	Memo       string
	AddressID  int `json:"address_id"`
	CouponID   int `json:"coupon_id"`
	Amount     float64
	Pay        float64
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}
