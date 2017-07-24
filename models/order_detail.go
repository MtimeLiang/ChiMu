package models

type OrderDetail struct {
	ID          int `json:"id"`
	OID         int `json:"oid"`
	PID         int `json:"pid"`
	Count       int
	Price       float64
	ProductInfo Product
}
