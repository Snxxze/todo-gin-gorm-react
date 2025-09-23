package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email		  string	`gorm:"email" json:"email"`
	Password  string	`gorm:"password" json:"-"`

	Todos []Todo	`json:"todos,omitempty"`
}