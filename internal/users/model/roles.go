package model

import "time"

type Roles struct {
	Id        uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"column:name;unique;not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
