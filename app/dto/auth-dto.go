package dto

import "github.com/google/uuid"

type LoginDto struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterDto struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min=1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginAdminResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	AccessToken string    `json:"accessToken"`
}

type LoginUserResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	AccessToken string    `json:"accessToken"`
}
