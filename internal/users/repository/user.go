package repository

import (
	"black-owl/internal/users/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	writeDb *gorm.DB
	readDb  *gorm.DB
}

func NewUserRepository(write, read *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		writeDb: write,
		readDb:  read,
	}
}
