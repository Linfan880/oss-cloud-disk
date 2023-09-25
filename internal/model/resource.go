package model

import "time"

type Resource struct {
	Id       int64     `gorm:"column:id;primary_key" json:"id"`
	Identity string    `gorm:"column:id;" json:"identity"` // 作为外键使用
	Hash     string    `gorm:"column:hash" json:"hash"`
	Name     string    `gorm:"column:name;" json:"name"`
	Ext      string    `gorm:"column:ext;" json:"ext"`
	Size     int64     `gorm:"column:size;" json:"size"`
	Path     string    `gorm:"column:path;" json:"path"`
	IsShared bool      `gorm:"column:is_shared" json:"is_shared"`
	CreateAt time.Time `gorm:"column:created_at;type:datatime" json:"create_at"`
	UpdateAt time.Time `gorm:"column:updated_at;type:datatime" json:"update_at"`
	DeleteAt time.Time `gorm:"column:deleted_at;type:datatime" json:"delete_at"`
}

func (table *Resource) TableName() string {
	return "resource"
}
