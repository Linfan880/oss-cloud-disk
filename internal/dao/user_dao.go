package dao

import (
	"errors"
	"file-system/internal/global"
	"file-system/internal/model"
	"file-system/internal/utils"
	"file-system/internal/web/types"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryInterface interface {
	CreateUser(req *types.UserRegisterReq) error
	UserLogin(req *types.UserLoginReq) (*model.User, error)
	DeleteUser(name string) error
	GetUserDetail(req *types.UserDetailReq) (*model.User, error)
}

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(req *types.UserRegisterReq) error {
	// 创建User对象
	user := &model.User{
		Name:     req.Username,
		Identity: utils.GenerateUuid(),
		Password: utils.Md5(req.Password), // 实际上你应该把密码加密后再存储，而不是直接存储密码
		Phone:    req.Phone,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	if err := global.MysqlConn.Debug().Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UserLogin(req *types.UserLoginReq) (*model.User, error) {
	user := &model.User{}
	if err := global.MysqlConn.Debug().Model(&model.User{}).Where("name = ?", req.Username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户未注册")
		}
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) DeleteUser(name string) error {
	if err := global.MysqlConn.Debug().Model(&model.User{}).Delete(&model.User{}, "name = ?", name).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserDetail(req *types.UserDetailReq) (*model.User, error) {
	user := &model.User{}
	if err := global.MysqlConn.Debug().Model(&model.User{}).Where("identity = ?", req.Identity).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// 检查用户是否存在
func CheckUserExist(name string) (bool, error) {
	var count int64
	if err := global.MysqlConn.Debug().Model(&model.User{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
