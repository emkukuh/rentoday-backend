package repository

import (
	"rentoday.id/app/database"
	"rentoday.id/app/model"
)

type _wardrobeRepository interface {
	InsertWardrobe(wardrobe model.Wardrobe) (model.Wardrobe, error)
	FindWardrobesByEmail(email string) ([]model.Wardrobe, error)
}

type wardrobeRepostory struct {}

var WardrobeRepository _wardrobeRepository = &wardrobeRepostory{}

func (w *wardrobeRepostory) InsertWardrobe(wardrobe model.Wardrobe) (model.Wardrobe, error) {
	var newWardrobe model.Wardrobe
	res := database.DB.Create(&newWardrobe)
	return newWardrobe, res.Error
}

func (w *wardrobeRepostory) FindWardrobesByEmail(email string) ([]model.Wardrobe, error) {
	return make([]model.Wardrobe, 0), nil
}