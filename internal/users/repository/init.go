package repository

import (
	"black-owl/internal/users/interfaces"
	"gorm.io/gorm"
)

type Init struct {
	UserRepo interfaces.UserRepository
}

func NewInit(write, read *gorm.DB) *Init {
	return &Init{
		UserRepo: NewUserRepository(write, read),
	}
}
