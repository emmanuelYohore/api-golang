package models

import "github.com/google/uuid"

type Book struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
}
