package model

import "github.com/google/uuid"

type WardrobeOwner struct {
	BaseModel
	UserID     uuid.UUID `gorm:"type:uuid" json:"-"`
	WardrobeID uuid.UUID `gorm:"type:uuid" json:"-"`
	Name       string    `json:"name"`
}
