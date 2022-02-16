package service

import (
	"rentoday.id/app/model"
	"rentoday.id/app/repository"
)

var ( 
	wardrobeCategoryRepo = repository.NewWardrobeCategory()
)

type ConfigService interface {
	GetWardrobeCategories() ([]model.WardrobeCategory, error)
}

type configService struct {}

func NewConfigService() ConfigService {
	return &configService{}
}

func (c *configService) GetWardrobeCategories() ([]model.WardrobeCategory, error) {
	categories, err := wardrobeCategoryRepo.FindAll()
	return categories, err
}