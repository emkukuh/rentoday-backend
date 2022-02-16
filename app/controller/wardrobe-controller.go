package controller

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/model"
	"rentoday.id/app/service"
)

var wardrobeService = service.NewWardrobeService()

type WardrobeController interface {
	FindAll() []model.Wardrobe
	Save(ctx  *gin.Context) model.Wardrobe
}

type wardrobeController struct {}

func NewWardrobeController() WardrobeController {
	return &wardrobeController{}
}

func (c *wardrobeController) FindAll() []model.Wardrobe {
	return wardrobeService.FindAll()
}

func (c *wardrobeController) Save(ctx *gin.Context) model.Wardrobe {
	var wardrobe model.Wardrobe
	ctx.BindJSON(&wardrobe)
	wardrobeService.Save(wardrobe)
	return wardrobe
}