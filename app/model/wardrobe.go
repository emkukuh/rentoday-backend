package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Wardrobe struct {
	BaseModel
	UserID       string         `json:"userId,omitempty"`
	LabelName    string         `json:"labelName"`
	Category     string         `json:"category"`
	Condition    string         `json:"condition"`
	LabelSize    string         `json:"labelSize"`
	SizeDetailID uuid.UUID      `json:"-"`
	SizeDetail   SizeDetail     `gorm:"foreignKey:SizeDetailID;OnDelete:CASCADE;" json:"sizeDetail"`
	Material     string         `json:"material"`
	Colors       pq.StringArray `gorm:"type:text[]" json:"colors"`
	Defects      pq.StringArray `gorm:"type:text[]" json:"defects"`
	// Images       string     `gorm:"type:text[]" json:"images"`
}
