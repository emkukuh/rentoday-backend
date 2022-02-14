package service

import (
	database "rentoday.id/app/database/model"
	"rentoday.id/app/model"
)

func getConfig() (*model.Configs, error) {
	var configs model.Configs
	database.DB.Find(&configs{})
}