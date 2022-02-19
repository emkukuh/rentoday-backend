package model

type WardrobeCategory struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
