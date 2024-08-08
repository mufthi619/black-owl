package model

import "time"

type User struct {
	Id         uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	RoleId     uint64     `json:"role_id"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Name       string     `json:"name"`
	LastAccess *time.Time `json:"lastAccess" gorm:"column:last_access;default:CURRENT_TIMESTAMP"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type RoleRight struct {
	Id        uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	RoleId    uint64     `json:"role_id"`
	Route     string     `json:"route"`
	Section   string     `json:"section"`
	Path      string     `json:"path"`
	RCreate   int        `json:"r_create"`
	RRead     int        `json:"r_read"`
	RUpdate   int        `json:"r_update"`
	RDelete   int        `json:"r_delete"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
}
