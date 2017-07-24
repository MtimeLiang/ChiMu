package models

type Promotion struct {
	ID          int `json:"id"`
	PID         int `json:"pid"`
	Name        string
	Description string
}
