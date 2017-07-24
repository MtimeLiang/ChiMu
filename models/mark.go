package models

type Mark struct {
	ID     int `json:"id"`
	PID    int `json:"pid"`
	Name   string
	Detail string
}
