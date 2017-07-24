package models

type Product struct {
	ID               int `json:"id"`
	Title            string
	SubMessage       string `json:"submessage"`
	Price            float64
	Volume           int
	Image            string
	Description      string
	OriginPrice      float64 `json:"origprice"`
	Count            int
	DescriptionImage string  `json:"description_image"`
	FreightMoney     float64 `json:"freight_money"`
	Point            int
	Sales            int
	Images           []string
	DescImages       []string
	Invalid          int // 默认下架
}
