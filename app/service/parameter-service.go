package service

import (
	"rentoday.id/app/model"
	"rentoday.id/app/repository"
)

var (
	wardrobeCategoryRepo = repository.ParameterRepository
)

type ParameterService interface {
	GetWardrobeCategories() ([]model.WardrobeCategory, error)
	GetWardrobeMaterials() ([]model.WardrobeMaterial, error)
}

type parameterService struct{}

func NewParameterService() ParameterService {
	return &parameterService{}
}

func (c *parameterService) GetWardrobeCategories() ([]model.WardrobeCategory, error) {
	categories, err := wardrobeCategoryRepo.FindAllWardrobeCategory()
	return categories, err
}

func (c *parameterService) GetWardrobeMaterials() ([]model.WardrobeMaterial, error) {
	materials, err := wardrobeCategoryRepo.FindAllWardrobeMaterial()
	return materials, err
}
