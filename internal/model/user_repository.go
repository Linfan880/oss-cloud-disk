package model

import (
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	Id           int64          `gorm:"column:id;primary_key" json:"id"`
	CreateAt     time.Time      `gorm:"column:created_at;type:datatime" json:"create_at"`
	UpdateAt     time.Time      `gorm:"column:updated_at;type:datatime" json:"update_at"`
	DeleteAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datatime" json:"delete_at"`
	UserName     string         `gorm:"column:user_name" json:"user_name"`
	ResourceName string         `gorm:"column:resource_name" json:"resource_name"`
	Ext          string         `gorm:"column:ext" json:"ext"`
	Size         int64          `gorm:"column:size;" json:"size"`
	Path         string         `gorm:"column:path;" json:"path"`
	IsShared     bool           `gorm:"column:is_shared" json:"is_shared"`
}

func (table *UserRepository) TableName() string {
	return "user_repository"
}
