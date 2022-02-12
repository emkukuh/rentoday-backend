package model

type Wardrobe struct {
	ID			string	`json:"id"`
	Name	 	string	`json:"name"`
	Size 		string  `json:"size"`
	Images		[]string  `json:"images"`
}