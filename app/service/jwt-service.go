package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"rentoday.id/app/config"
	"rentoday.id/app/constant"
)

type JwtServiceInterface interface {
	GenerateToken(email string) string 
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtCustomClaim struct {
	Email				 string `json:"email"`
	jwt.StandardClaims
}

type JwtService struct {
	SecretKey 	string
	Issuer 		string
}

func NewJwtService() JwtServiceInterface {
	return &JwtService{
		SecretKey: genereateSecretKey(),
		Issuer: config.JwtIssuer(),
	}
}

func genereateSecretKey() string {
	secretKey := config.JwtSecretKey()
	if len(secretKey) == 0 {
		panic("secret key is empty")
	}
	return secretKey
}

func (s *JwtService) GenerateToken(email string) string {
	claims := &JwtCustomClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: constant.OneDay24Hour,
			Issuer: s.Issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func (t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(s.SecretKey), nil
	
	})
}