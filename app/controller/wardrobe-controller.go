package controller

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/model"
	"rentoday.id/app/service"
)

type WardrobeControllerInterface interface {
	FindAll() []model.Wardrobe
	Save(ctx  *gin.Context) model.Wardrobe
}

type WardrobeController struct {
	service service.WardrobeServiceInterface
}

func New(service service.WardrobeServiceInterface) WardrobeController {
	return WardrobeController {
		service: service,
	}
}

func (c *WardrobeController) FindAll() []model.Wardrobe {
	return c.service.FindAll()
}

func (c *WardrobeController) Save(ctx *gin.Context) model.Wardrobe {
	var wardrobe model.Wardrobe
	ctx.BindJSON(&wardrobe)
	c.service.Save(wardrobe)
	return wardrobe
}