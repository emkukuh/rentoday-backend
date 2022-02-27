package model

import "github.com/google/uuid"

type WardrobeOwner struct {
	BaseModel
	UserID     uuid.UUID
	WardrobeID uuid.UUID
}
