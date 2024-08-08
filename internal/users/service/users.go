package service

import (
	"black-owl/internal/users/interfaces"
	"gorm.io/gorm"
)

type User struct {
	writeDb *gorm.DB
	readDb  *gorm.DB
}

func NewUserService(write, read *gorm.DB) interfaces.UserService {
	return &User{
		writeDb: write,
		readDb:  read,
	}
}
