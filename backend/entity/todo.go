package entity

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `gorm:"title" json:"title"`
	Description string `gorm:"description" json"description,omitempty"`
	Status		string `gorm:"status" json:"status"`

	UserID	uint `json:"userId"`
	User User `json:"-"`
}