package service

import "rentoday.id/app/model"


type WardrobeService interface {
	Save(model.Wardrobe) model.Wardrobe
	FindAll() []model.Wardrobe
}

type wardrobeService struct {
	wardrobes []model.Wardrobe
}

func NewWardrobeService() WardrobeService {
	return &wardrobeService{}
}

func (service *wardrobeService) Save(video model.Wardrobe) model.Wardrobe {
	service.wardrobes = append(service.wardrobes, video)
	return video
}

func (service *wardrobeService) FindAll() []model.Wardrobe {
	return service.wardrobes
}