package model

type Wardrobe struct {
	BaseModel
	OwnerID			User		`gorm:"foreignKey:ID;references:ID" json:"ownerId"`
	LabelName		string		`json:"labelName"`
	Category    	string  	`json:"category"`
	Condition		string		`json:"condition"`
	LabelSize 		string 		`json:"labelSize"`
	// SizeDetailID	SizeDetail	`json:"sizeDetailId"`
	Material		string		`json:"material"`
	Colors			string		`gorm:"type:text[]" json:"colors"`
	Defect 			string 		`json:"defect"`
	Images			string 		`gorm:"type:text[]" json:"images"`
}