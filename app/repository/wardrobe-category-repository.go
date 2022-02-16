package repository

import (
	"rentoday.id/app/database"
	"rentoday.id/app/model"
)

type WardrobeCategory interface {
	FindAll() ([]model.WardrobeCategory, error)
}

type wardrobeCategory struct {}

func NewWardrobeCategory() WardrobeCategory {
	return &wardrobeCategory{}
}

func (w *wardrobeCategory) FindAll() ([]model.WardrobeCategory, error) {
	var category []model.WardrobeCategory
	res := database.DB.Select("*").Find(&category)
	return category, res.Error
}