package model

type User struct {
	ID    		uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name 		string	`json:"name"`
	Email 		string	`json:"email"`
	Password	string	`gorm:"->;<-;not null" json:"-"`
	Token		string 	`gorm:"-" json:"token,omitempty"`
}