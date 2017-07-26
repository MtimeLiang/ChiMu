package models

type Type struct {
	ID   int `json:"id" orm:"column(id)"`
	PID  int `json:"pid" orm:"column(pid)"`
	Name string
	Type int
}
