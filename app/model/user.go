package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 	string	`gorm:"primaryKey" json:"name"`
	Email 	string	`json:"email"`
}