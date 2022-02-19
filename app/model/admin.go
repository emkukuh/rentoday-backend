package model

type Admin struct {
	BaseModel
	Name 		string		`json:"name"`
	Email 		string		`json:"email"`
	Password	string		`gorm:"->;<-;not null" json:"-"`
	AccessToken	string 		`gorm:"-" json:"accessToken,omitempty"`
}