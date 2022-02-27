package service

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"rentoday.id/app/dto"
	"rentoday.id/app/model"
	"rentoday.id/app/repository"
)

var wardrobeRepo = repository.WardrobeRepository

type WardrobeService interface {
	Create(wardrobe dto.AddWardrobeRequest) (model.Wardrobe, error)
	FindAll() ([]model.Wardrobe, error)
	GetListByUserId(userId string) ([]model.Wardrobe, error)
}

type wardrobeService struct{}

func NewWardrobeService() WardrobeService {
	return &wardrobeService{}
}

func (service *wardrobeService) Create(wardrobe dto.AddWardrobeRequest) (model.Wardrobe, error) {
	newWardrobe := model.Wardrobe{}
	err := smapping.FillStruct(&newWardrobe, smapping.MapFields(&wardrobe))
	if err != nil {
		log.Printf("failed to map %v", err)
		return newWardrobe, errors.New(err.Error())
	}
	log.Print("newWardrobe = ", newWardrobe)
	wardrobeRes, err := wardrobeRepo.InsertWardrobe(newWardrobe)
	return wardrobeRes, err
}

func (service *wardrobeService) FindAll() ([]model.Wardrobe, error) {
	return wardrobeRepo.FindAllWardrobe()
}

func (service *wardrobeService) GetListByUserId(userId string) ([]model.Wardrobe, error) {
	return wardrobeRepo.FindAllWardrobeByUserId(userId)
}
