package types

type HTTPResp struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

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
