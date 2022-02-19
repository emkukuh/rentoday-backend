package model

type WardrobeMaterial struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
