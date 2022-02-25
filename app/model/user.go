package model

type User struct {
	BaseModel
	Name          string          `gorm:"index" json:"name"`
	Email         string          `gorm:"index" json:"email"`
	PhoneNumber   string          `json:"phoneNumber"`
	Password      string          `gorm:"->;<-;not null" json:"-"`
	WardrobeOwner []WardrobeOwner `gorm:"foreignKey:UserID"`
	AccessToken   string          `gorm:"-" json:"accessToken,omitempty"`
}
