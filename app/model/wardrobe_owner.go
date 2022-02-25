package model

import "github.com/google/uuid"

type WardrobeOwner struct {
	BaseModel
	UserID     uuid.UUID
	WardrobeID uuid.UUID
	// WardrobeID uuid.UUID `gorm:"type:uuid" json:"-"`
	// Name string `json:"name"`
}
