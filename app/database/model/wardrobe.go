package database

type Wardrobe struct {
	ID					string	`gorm:"primaryKey"`
	Name			 	string
	WardrobeCategoryID  string
	WardrobeMaterialID	string
	Size 				string
	Images				string
}