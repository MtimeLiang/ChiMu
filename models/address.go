package models

type Address struct {
	ID         int
	UID        int
	Name       string
	Tel        string
	Address    string
	IsDefault  int `json:"is_default"`
	IsSelected int `json:"is_selected"`
	IsDelete   int `json:"is_delete"`
}
