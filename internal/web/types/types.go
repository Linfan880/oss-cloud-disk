package types

import "time"

type HTTPResp struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

// user相关
type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

type UserDeleteReq struct {
	Username string `json:"username"`
}

type UserDetailReq struct {
	Identity string `json:"identity"`
}

type UserDetailResp struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

// 资源相关
type ResourceUploadReq struct {
	Username       string `json:"username"`
	Identity       string `json:"identity,omitempty"`
	Hash           string `json:"hash,optional"`
	Name           string `json:"name,optional"`
	Ext            string `json:"ext,optional"`
	Size           int64  `json:"size,optional"`
	ParentIdentity string `json:"parent_identity"`
	Path           string `json:"path,optional"`
}

type ResourceUploadResp struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	Path     string `json:"path,optional"`
}

type ResourceDeleteReq struct {
	Identity string `json:"identity"`
}

type ResourceDeleteResp struct {
}

type ResourceDetailReq struct {
	Identity string `json:"identity"`
}

type ResourceDetailResp struct {
	Name     string    `json:"name"`
	Ext      string    `json:"ext"`
	Size     int64     `json:"size"`
	Path     string    `json:"path"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type ShareResourceReq struct {
	Username string `json:"username"`
	Identity string `json:"identity"`
}

type ShareResourceResp struct {
}

// 用户-资源相关
type UserResourceListReq struct {
	Identity string `json:"identity,optional"`
	Page     int    `json:"page,optional"`
	Size     int    `json:"size,optional"`
}

type UserResourceListResp struct {
	Count int64           `json:"count"`
	List  []*UserResource `json:"list"`
}

type UserResource struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
}
