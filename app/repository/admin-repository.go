package repository

import (
	"rentoday.id/app/database"
	"rentoday.id/app/model"
)

type _adminRepository interface {
	InsertUser(user model.Admin) (model.Admin, error)
	VerifyCredential(email string, password string) (model.Admin, error)
	FindUserByEmail(email string) (model.Admin, error)
}

type adminRepository struct {}

var AdminRepository	_adminRepository = &adminRepository{}

func (r *adminRepository) InsertUser(admin model.Admin) (model.Admin, error) {
	res := database.DB.Create(&admin)
	return admin, res.Error
} 

func (r *adminRepository) VerifyCredential(email string, password string) (model.Admin, error) {
	var admin model.Admin
	res := database.DB.Where("email=?", email).Take(&admin)
	return admin, res.Error
}

func (r *adminRepository) FindUserByEmail(email string) (model.Admin, error) {
	var admin model.Admin
	res := database.DB.Where("email=?", email).Take(&admin)
	return admin, res.Error
}