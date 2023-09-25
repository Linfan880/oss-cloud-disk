package repository

import (
	"errors"
	"file-system/internal/global"
	"file-system/internal/model"
	"file-system/internal/utils"
	"time"
)

type UserRepositoryInterface interface {
	CreateUser(name, password, phone string) error
	DeleteUser(name string) error
}

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(name, password, phone string) error {
	// 检查用户是否已存在
	exist, err := utils.CheckUserExist(name)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("user already exists")
	}

	// 创建User对象
	user := &model.User{
		Name:         name,
		Identity:     utils.GenerateUuid(),
		PasswordHash: utils.Md5(password), // 实际上你应该把密码加密后再存储，而不是直接存储密码
		Phone:        phone,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
	}

	if err := global.MysqlConn.Debug().Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteUser(name string) error {
	// 检查用户是否存在
	exist, err := utils.CheckUserExist(name)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("user does not exist")
	}

	if err := global.MysqlConn.Debug().Delete(&model.User{}, "name = ?", name).Error; err != nil {
		return err
	}

	return nil
}
