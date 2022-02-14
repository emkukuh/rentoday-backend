package database

type WardrobeMaterial struct {
	ID			string	`gorm:"primaryKey"`
	Name	 	string	`json:"name"`
}
