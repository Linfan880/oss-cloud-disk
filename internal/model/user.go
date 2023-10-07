package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int64          `gorm:"column:id;primary_key" json:"id"`
	Identity string         `gorm:"column:identity" json:"identity"` //作为外键使用
	Name     string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Password string         `gorm:"column:password;type:varchar(255)" json:"password"`
	Phone    string         `gorm:"column:phone;type:varchar(11)" json:"phone"`
	CreateAt time.Time      `gorm:"column:created_at;type:datatime(3)" json:"create_at"`
	UpdateAt time.Time      `gorm:"column:updated_at;type:datatime(3)" json:"update_at"`
	DeleteAt gorm.DeletedAt `gorm:"column:deleted_at;type:datatime(3)" json:"delete_at"`
}

func (table *User) TableName() string {
	return "user"
}
