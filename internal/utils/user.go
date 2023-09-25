package utils

import (
	"file-system/internal/global"
	"file-system/internal/model"
)

// 检查用户是否存在
func CheckUserExist(name string) (bool, error) {
	var count int64
	if err := global.MysqlConn.Model(&model.User{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
