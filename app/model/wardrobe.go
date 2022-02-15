package model

type Wardrobe struct {
	ID			string		`gorm:"primaryKey" json:"id"`
	Name	 	string		`json:"name"`
	Category    string  	`json:"category"`
	Material	string		`json:"material"`
	Size 		string 		`json:"size"`
	Images		[]string 	`gorm:"type:text" json:"images"`
}