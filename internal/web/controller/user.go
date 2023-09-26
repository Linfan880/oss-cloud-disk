package controller

import (
	"file-system/internal/dao"
	"file-system/internal/utils"
	"file-system/internal/web/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	d := dao.NewUserRepository()
	var userInfo types.UserRegisterReq
	// 将请求的Form表单数据绑定到form结构体中
	if ctx.ShouldBind(&userInfo) == nil {
		// 如果绑定成功对form进行处理
	}

	// 检查用户是否已存在
	exist, err := utils.CheckUserExist(userInfo.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.HTTPResp{
			ErrCode: http.StatusInternalServerError,
			ErrMsg:  err.Error(),
			Data:    nil,
		})
		return
	}
	if exist {
		ctx.JSON(http.StatusInternalServerError, types.HTTPResp{
			ErrCode: http.StatusInternalServerError,
			ErrMsg:  "用户已存在",
			Data:    nil,
		})
		return
	}

	if err = d.CreateUser(&userInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, types.HTTPResp{
			ErrCode: http.StatusInternalServerError,
			ErrMsg:  "用户注册失败",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, "创建用户成功")
}

func UserLogin(ctx *gin.Context) {
	d := dao.NewUserRepository()
	var userReq types.UserLoginReq

	if ctx.ShouldBind(&userReq) == nil {
		// 如果绑定成功对form进行处理
		resp, err := d.UserLogin(&userReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.HTTPResp{
				ErrCode: http.StatusInternalServerError,
				ErrMsg:  "用户登录绑定信息失败",
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusOK, types.HTTPResp{
			ErrCode: http.StatusOK,
			ErrMsg:  "用户登录成功",
			Data:    resp,
		})
		return
	}
}

func DeleteUser(ctx *gin.Context) {
	// 检查用户是否存在
	//exist, err := utils.CheckUserExist(name)
	//if err != nil {
	//	return err
	//}
	//if !exist {
	//	return errors.New("user does not exist")
	//}
}
