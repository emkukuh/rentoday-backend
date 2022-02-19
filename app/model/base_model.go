package model

import "time"

type BaseModel struct {
	ID        uint64 `gorm:"primarykey:auto_increment" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
