package repository

import (
	"rentoday.id/app/database"
	"rentoday.id/app/model"
)

type _wardrobeRepository interface {
	InsertWardrobe(wardrobe model.Wardrobe) (model.Wardrobe, error)
	FindAllWardrobe() ([]model.Wardrobe, error)
	FindWardrobesByEmail(email string) ([]model.Wardrobe, error)
}

type wardrobeRepostory struct{}

var WardrobeRepository _wardrobeRepository = &wardrobeRepostory{}

func (w *wardrobeRepostory) InsertWardrobe(wardrobe model.Wardrobe) (model.Wardrobe, error) {
	res := database.DB.Create(&wardrobe)
	return wardrobe, res.Error
}

func (w *wardrobeRepostory) FindAllWardrobe() ([]model.Wardrobe, error) {
	var wardrobes []model.Wardrobe
	res := database.DB.Find(&wardrobes)
	return wardrobes, res.Error
}

func (w *wardrobeRepostory) FindWardrobesByEmail(email string) ([]model.Wardrobe, error) {
	return make([]model.Wardrobe, 0), nil
}
