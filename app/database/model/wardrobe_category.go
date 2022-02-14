package database

type WardrobeCategory struct {
	ID			string	`gorm:"primaryKey"`
	Name	 	string
}