package models

type Mark struct {
	ID     int `json:"id" orm:"column(id)"`
	PID    int `json:"pid" orm:"column(pid)"`
	Name   string
	Detail string
}
