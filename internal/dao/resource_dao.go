package dao

import (
	"file-system/internal/global"
	"file-system/internal/model"
	"file-system/internal/utils"
	"file-system/internal/web/types"
)

type ResourceRepositoryInterface interface {
	UploadResource(*types.ResourceUploadReq) error
	DownloadResource() error
	GetResourceDetail(*types.ResourceDetailReq) (*types.ResourceDetailResp, error)
	ShareResource(*types.ShareResourceReq) error
	DeleteResource(*types.ResourceDeleteReq) error
}

type ResourceRepository struct {
}

func NewResourceRepository() ResourceRepositoryInterface {
	return &ResourceRepository{}
}

func (r *ResourceRepository) UploadResource(req *types.ResourceUploadReq) error {
	// 资源池内加入资源
	re := model.Resource{
		Identity:       utils.GenerateUuid(),
		Hash:           utils.Md5(req.Name),
		Name:           req.Name,
		Ext:            req.Ext,
		Size:           req.Size,
		ParentIdentity: req.ParentIdentity,
		Path:           req.Path,
		IsShared:       false,
	}

	if err := global.MysqlConn.Debug().Model(&model.Resource{}).Create(re).Error; err != nil {
		return err
	}

	// 资源所属用户关系表也要插入
	ur := model.UserRepository{
		UserName:     req.Username,
		ResourceName: req.Name,
		Ext:          req.Ext,
		Size:         req.Size,
		Path:         req.Path,
		IsShared:     false,
	}
	if err := global.MysqlConn.Debug().Model(&model.UserRepository{}).Create(ur).Error; err != nil {
		return err
	}

	return nil
}
func (r *ResourceRepository) DownloadResource() error {
	return nil
}
func (r *ResourceRepository) GetResourceDetail(req *types.ResourceDetailReq) (*types.ResourceDetailResp, error) {
	re := &model.Resource{}
	if err := global.MysqlConn.Debug().Model(&model.Resource{}).Where("identity = ?", req.Identity).First(re).Error; err != nil {
		return nil, err
	}

	resp := &types.ResourceDetailResp{
		Name:     re.Name,
		Ext:      re.Ext,
		Size:     re.Size,
		Path:     re.Path,
		CreateAt: re.CreateAt,
		UpdateAt: re.UpdateAt,
	}

	return resp, nil
}
func (r *ResourceRepository) ShareResource(req *types.ShareResourceReq) error {
	re := &model.Resource{}
	if err := global.MysqlConn.Debug().Model(&model.Resource{}).Where("identity = ?", req.Identity).First(re).Error; err != nil {
		return err
	}

	u := &model.User{}
	if err := global.MysqlConn.Debug().Model(&model.User{}).Where("username = ?", req.Username).First(u).Error; err != nil {
		return err
	}

	if err := global.MysqlConn.Debug().Model(&model.Resource{}).Where("identity = ?", req.Identity).Update("is_shared", "true").Error; err != nil {
		return err
	}

	sr := model.SharedPool{
		UserIdentity:     u.Identity,
		UserName:         u.Name,
		ResourceIdentity: req.Identity,
		ResourceName:     re.Name,
		Ext:              re.Ext,
		Size:             re.Size,
		Path:             re.Path,
		ClickNum:         0,
	}
	if err := global.MysqlConn.Debug().Model(&model.UserRepository{}).Create(sr).Error; err != nil {
		return err
	}

	return nil
}

func (r *ResourceRepository) DeleteResource(req *types.ResourceDeleteReq) error {
	if err := global.MysqlConn.Debug().Model(&model.Resource{}).Delete(&model.User{}, "identity = ?", req.Identity).Error; err != nil {
		return err
	}

	return nil
}

// 检查资源是否存在
func CheckResourceExist(name string) (bool, error) {
	var count int64
	if err := global.MysqlConn.Debug().Model(&model.Resource{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
