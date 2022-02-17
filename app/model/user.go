package model

type User struct {
	BaseModel
	ID    		uint64  `gorm:"primaryKey:autoIncrement" json:"id"`
	Name 		string	`gorm:"index" json:"name"`
	Email 		string	`json:"email"`
	PhoneNumber	string	`json:"phoneNumber"`
	Password	string	`gorm:"->;<-;not null" json:"-"`
	AccessToken	string 	`gorm:"-" json:"accessToken,omitempty"`
}