package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"rentoday.id/app/dto"
	"rentoday.id/app/model"
	"rentoday.id/app/repository"
)

var (
	userRepo = repository.NewUserRepository()
	jwtService = NewJwtService()
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	Register(user dto.RegisterDto) (model.User, error)
}

type authService struct {}

func NewAuthService() AuthService {
	return &authService{}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := userRepo.VerifyCredential(email, password)
	if v, ok := res.(model.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) Register(user dto.RegisterDto) (model.User, error) {
	newUser := model.User{}
	isDuplicated, errIsDuplicated := isEmailDuplicated(user.Email)
	if errIsDuplicated != nil {
		return newUser, errIsDuplicated
	}
	if isDuplicated {
		return newUser, errors.New("email sudah terdaftar")
	}
	errSmapping := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if errSmapping != nil {
		log.Fatalf("failed to map %v", errSmapping)
	}
	token := jwtService.GenerateToken(strconv.FormatUint(newUser.ID, 10))
	newUser.Token = token
	res, err := userRepo.InsertUser(newUser)
	return res, err
}

func isEmailDuplicated(email string) (bool, error) {
	_, err := userRepo.FindUserByEmail(email) 
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func comparePassword(hashPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}