package config

import "os"

func JwtSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func JwtIssuer() string {
	return "Rentoday"
}