package model

import "github.com/google/uuid"

type Wardrobe struct {
	BaseModel
	UserID       uuid.UUID  `gorm:"type:uuid" json:"-"`
	LabelName    string     `json:"labelName"`
	Category     string     `json:"category"`
	Condition    string     `json:"condition"`
	LabelSize    string     `json:"labelSize"`
	SizeDetailID uuid.UUID  `gorm:"type:uuid" json:"sizeDetailId"`
	SizeDetail   SizeDetail `json:"sizeDetail"`
	Material     string     `json:"material"`
	// Colors       string     `gorm:"type:text[]" json:"colors"`
	// Defects string `gorm:"type:text[]" json:"defects"`
	// Images       string     `gorm:"type:text[]" json:"images"`
}
