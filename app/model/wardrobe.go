package model

type Wardrobe struct {
	BaseModel
	UserID       uint
	LabelName    string     `json:"labelName"`
	Category     string     `json:"category"`
	Condition    string     `json:"condition"`
	LabelSize    string     `json:"labelSize"`
	SizeDetailID uint       `json:"sizeDetailId"`
	SizeDetail   SizeDetail `json:"sizeDetail"`
	Material     string     `json:"material"`
	Colors       string     `gorm:"type:text[]" json:"colors"`
	Defects      string     `json:"defects[]"`
	Images       string     `gorm:"type:text[]" json:"images"`
}
