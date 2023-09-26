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
	UserLogin(req *types.UserLoginReq) (*types.UserLoginResp, error)
	DeleteUser(name string) error
}

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(req *types.UserRegisterReq) error {
	// 创建User对象
	user := &model.User{
		Name:         req.Username,
		Identity:     utils.GenerateUuid(),
		PasswordHash: utils.Md5(req.Password), // 实际上你应该把密码加密后再存储，而不是直接存储密码
		Phone:        req.Phone,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
	}

	if err := global.MysqlConn.Debug().Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UserLogin(req *types.UserLoginReq) (*types.UserLoginResp, error) {
	user := &model.User{}
	if err := global.MysqlConn.Debug().Where("name = ?", req.Username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户未注册")
		}
		return nil, err
	}

	// 验证密码是否一致
	if user.PasswordHash != utils.Md5(req.Password) {
		return nil, errors.New("密码错误")
	}
	// 2、生成token
	token, err := utils.GenerateToken(user.Identity, user.Name, global.TokenExpire)
	if err != nil {
		return nil, err
	}

	resp := new(types.UserLoginResp)
	resp.Token = token

	return resp, nil
}

func (u *UserRepository) DeleteUser(name string) error {
	if err := global.MysqlConn.Debug().Delete(&model.User{}, "name = ?", name).Error; err != nil {
		return err
	}

	return nil
}
