package repository

import (
	"rentoday.id/app/database"
	"rentoday.id/app/model"
)

type UserRepository interface {
	InsertUser(user model.User) (model.User, error)
	VerifyCredential(email string, password string) (model.User, error)
	FindUserByEmail(email string) (model.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}

}
func (r *userRepository) InsertUser(user model.User) (model.User, error) {
	res := database.DB.Create(&user)
	return user, res.Error
}

func (r *userRepository) VerifyCredential(email string, password string) (model.User, error) {
	var user model.User
	res := database.DB.Where("email=?", email).Take(&user)
	return user, res.Error
}

func (r *userRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	res := database.DB.Where("email=?", email).Take(&user)
	return user, res.Error
}
