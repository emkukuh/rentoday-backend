package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"rentoday.id/app/config"
	"rentoday.id/app/constant"
)

type JwtServiceInterface interface {
	GenerateToken(userID uuid.UUID) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIdByToken(token string) (string, error)
}

type JwtCustomClaim struct {
	UserId uuid.UUID `json:"userId"`
	jwt.StandardClaims
}

type JwtService struct {
	SecretKey string
	Issuer    string
}

func NewJwtService() JwtServiceInterface {
	return &JwtService{
		SecretKey: genereateSecretKey(),
		Issuer:    config.JwtIssuer(),
	}
}

func genereateSecretKey() string {
	secretKey := config.JwtSecretKey()
	if len(secretKey) == 0 {
		panic("secret key is empty")
	}
	return secretKey
}

func (s *JwtService) GenerateToken(userId uuid.UUID) string {
	claims := &JwtCustomClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: constant.OneDay24Hour,
			Issuer:    s.Issuer,
			IssuedAt:  time.Now().Unix(),
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
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(s.SecretKey), nil

	})
}

func (s *JwtService) GetUserIdByToken(token string) (string, error) {
	currToken, err := s.ValidateToken(token)
	if err != nil {
		return "", err
	}
	claims := currToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userId"])
	return id, nil
}
