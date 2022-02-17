package model

type Wardrobe struct {
	BaseModel
	OwnerID			User		`json:"ownerId"`
	LabelName		string		`json:"name"`
	Category    	string  	`json:"category"`
	Condition		string		`json:"condition"`
	LabelSize 		string 		`json:"labelSize"`
	SizeDetailID	SizeDetail	`json:"sizeDetailId"`
	Material		string		`json:"material"`
	Colors			string		`gorm:"type:text[]" json:"colors"`
	Defect 			string 		`json:"defect"`
	Images			string 		`gorm:"type:text[]" json:"images"`
}