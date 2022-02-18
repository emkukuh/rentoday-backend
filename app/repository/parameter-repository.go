package repository

import (
	"rentoday.id/app/database"
	"rentoday.id/app/model"
)

type _parameterRepository interface {
	FindAllWardrobeCategory() ([]model.WardrobeCategory, error)
	FindAllWardrobeMaterial() ([]model.WardrobeMaterial, error)
}
type parameterRepository struct {}

var ParameterRepository _parameterRepository = &parameterRepository{}

func (p *parameterRepository) FindAllWardrobeCategory() ([]model.WardrobeCategory, error) {
	var categories []model.WardrobeCategory
	res := database.DB.Select("*").Find(&categories)
	return categories, res.Error
}

func (p *parameterRepository) FindAllWardrobeMaterial() ([]model.WardrobeMaterial, error) {
	var materials []model.WardrobeMaterial
	res := database.DB.Select("*").Find(&materials)
	return materials, res.Error
}