package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
	"rentoday.id/app/dto"
	"rentoday.id/app/model"
	"rentoday.id/app/repository"
)

type _authAdminService interface {
	VerifyCredential(email string, password string) (model.Admin, error)
	Register(user dto.RegisterDto) (model.Admin, error)
}

type authAdminService struct{}

var (
	AuthAdminService _authAdminService = &authAdminService{}
	adminRepo                          = repository.AdminRepository
)

func (a *authAdminService) VerifyCredential(email string, password string) (model.Admin, error) {
	admin, err := adminRepo.VerifyCredential(email, password)
	if err != nil {
		return admin, err
	}
	isComparedPasswordMatch := comparePassword(admin.Password, password)
	if admin.Email == email && isComparedPasswordMatch {
		return admin, nil
	}
	return admin, err
}

func (a *authAdminService) Register(user dto.RegisterDto) (model.Admin, error) {
	newAdmin := model.Admin{}
	isDuplicated, errIsDuplicated := isAdminEmailDuplicated(user.Email)
	if errIsDuplicated != nil {
		return newAdmin, errIsDuplicated
	}
	if isDuplicated {
		return newAdmin, errors.New("email sudah terdaftar")
	}
	err := smapping.FillStruct(&newAdmin, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed to map %v", err)
	}
	token := jwtService.GenerateToken(strconv.FormatUint(newAdmin.ID, 10))
	newAdmin.AccessToken = token
	newAdmin.Password = hashAndSaltPassword(user.Password)
	res, err := adminRepo.InsertUser(newAdmin)
	return res, err

}

func isAdminEmailDuplicated(email string) (bool, error) {
	_, err := adminRepo.FindUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}
