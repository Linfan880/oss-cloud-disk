package model

import (
	"gorm.io/gorm"
	"time"
)

type SharedPool struct {
	Id               int64          `gorm:"column:id;primary_key" json:"id"`
	UserIdentity     string         `gorm:"column:user_identity" json:"user-identity"`
	UserName         string         `gorm:"column:user_name" json:"user_name"`
	ResourceName     string         `gorm:"column:resource_name" json:"resource_name"`
	ResourceIdentity string         `gorm:"column:resource_identity" json:"resource_identity"`
	Ext              string         `gorm:"column:ext" json:"ext"`
	Size             int64          `gorm:"column:size;" json:"size"`
	Path             string         `gorm:"column:path;" json:"path"`
	ClickNum         int64          `gorm:"column:click_num" json:"click_num"`
	CreateAt         time.Time      `gorm:"column:created_at;type:datatime" json:"create_at"`
	UpdateAt         time.Time      `gorm:"column:updated_at;type:datatime" json:"update_at"`
	DeleteAt         gorm.DeletedAt `gorm:"column:deleted_at;type:datatime" json:"delete_at"`
}

func (table *SharedPool) TableName() string {
	return "shared_pool"
}
