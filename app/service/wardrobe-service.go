package service

import "rentoday.id/app/model"


type WardrobeServiceInterface interface {
	Save(model.Wardrobe) model.Wardrobe
	FindAll() []model.Wardrobe
}

type WardrobeService struct {
	wardrobes []model.Wardrobe
}

func New() WardrobeServiceInterface {
	return &WardrobeService{}
}

func (service *WardrobeService) Save(video model.Wardrobe) model.Wardrobe {
	service.wardrobes = append(service.wardrobes, video)
	return video
}

func (service *WardrobeService) FindAll() []model.Wardrobe {
	return service.wardrobes
}