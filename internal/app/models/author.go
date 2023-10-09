package models

type Author struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name"`
}
