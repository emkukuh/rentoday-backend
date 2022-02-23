package model

type User struct {
	BaseModel
	Name        string `gorm:"index" json:"name"`
	Wardrobes   []Wardrobe
	Email       string `gorm:"index" json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `gorm:"->;<-;not null" json:"-"`
	AccessToken string `gorm:"-" json:"accessToken,omitempty"`
}
