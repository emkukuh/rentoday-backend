package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

var jwtService = NewJwtService()

func comparePassword(hashPwd string, plainPwd string) bool {
	byteHash := []byte(hashPwd)
	bytePass := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
