package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title        string `json:"title" form:"title"`
	Author       string `json:"author" form:"author"`
	Published_at string `json:"published_at" form:"published_at"`
}
