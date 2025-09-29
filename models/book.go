package models

type Book struct {
	ID          string `gorm:"default:uuid_generate_v4()"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
