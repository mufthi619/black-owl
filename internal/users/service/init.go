package service

import (
	"black-owl/internal/users/interfaces"
	"gorm.io/gorm"
)

type Init struct {
	UserService interfaces.UserService
}

func NewInit(write, read *gorm.DB) *Init {
	return &Init{
		UserService: NewUserService(write, read),
	}
}
