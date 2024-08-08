package app

import (
	"black-owl/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseClient(config config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=dissable TimeZone=Asia/Jakarta",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
