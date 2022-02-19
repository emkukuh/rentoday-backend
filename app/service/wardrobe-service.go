package service

import (
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
}

type wardrobeService struct{}

func NewWardrobeService() WardrobeService {
	return &wardrobeService{}
}

func (service *wardrobeService) Create(wardrobe dto.AddWardrobeRequest) (model.Wardrobe, error) {
	newWardrobe := model.Wardrobe{}
	log.Print("++++++++++")
	err := smapping.FillStruct(&newWardrobe, smapping.MapFields(&wardrobe))
	log.Print("========")
	if err != nil {
		log.Fatalf("failed to map %v", err)
		return newWardrobe, err
	}
	wardrobeRes, err := wardrobeRepo.InsertWardrobe(newWardrobe)
	return wardrobeRes, err
}

func (service *wardrobeService) FindAll() ([]model.Wardrobe, error) {
	return wardrobeRepo.FindAllWardrobe()
}
