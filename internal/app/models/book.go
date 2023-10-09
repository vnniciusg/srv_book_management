package models

type Book struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `json:"title"`
	Author      uint   `json:"author"`
	PublishYear int    `json:"publish_year"`
}
