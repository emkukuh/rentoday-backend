package model

type Wardrobe struct {
	ID			string	`json:"id"`
	Name	 	string	`json:"name"`
	Category    string  `json:"category"`
	Material	string	`json:"material"`
	Size 		string  `json:"size"`
	Images		[]string  `json:"images"`
}