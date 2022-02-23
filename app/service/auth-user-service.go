package service

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"rentoday.id/app/dto"
	"rentoday.id/app/model"
	"rentoday.id/app/repository"
)

type _authUserService interface {
	VerifyCredential(email string, password string) interface{}
	Register(user dto.RegisterDto) (model.User, error)
}

type authUserService struct{}

var (
	AuthUserService _authUserService = &authUserService{}
	userRepo                         = repository.NewUserRepository()
)

func (service *authUserService) VerifyCredential(email string, password string) interface{} {
	user, err := userRepo.VerifyCredential(email, password)
	if err != nil {
		return false
	}
	responseDto := dto.LoginAdminResponseDto{}
	comparedPassword := comparePassword(user.Password, password)
	if user.Email == email && comparedPassword {
		err := smapping.FillStruct(&responseDto, smapping.MapFields(&user))
		if err != nil {
			return false
		}
		responseDto.ID = user.ID
		return responseDto
	}
	return false
}

func (service *authUserService) Register(user dto.RegisterDto) (model.User, error) {
	newUser := model.User{}
	isDuplicated, errIsDuplicated := isUserEmailDuplicated(user.Email)
	if errIsDuplicated != nil {
		return newUser, errIsDuplicated
	}
	if isDuplicated {
		return newUser, errors.New("email sudah terdaftar")
	}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed to map %v", err)
	}
	token := jwtService.GenerateToken(newUser.ID)
	newUser.AccessToken = token
	newUser.Password = hashAndSaltPassword(user.Password)
	res, err := userRepo.InsertUser(newUser)
	return res, err
}

func isUserEmailDuplicated(email string) (bool, error) {
	_, err := userRepo.FindUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func hashAndSaltPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("failed to hash password")
	}
	return string(hash)
}
