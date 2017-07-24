package models

type Banner struct {
	ID     int
	Imgurl string
	URL    string `json:"url"`
	Title  string
}
