package controller

import (
	"file-system/internal/dao"
	"file-system/internal/global"
	"file-system/internal/utils"
	"file-system/internal/web/tools"
	"file-system/internal/web/types"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	d := dao.NewUserRepository()
	var userInfo types.UserRegisterReq
	// 将请求的Form表单数据绑定到form结构体中
	if err := ctx.ShouldBind(&userInfo); err != nil {
		tools.StatusInternalErr(ctx, "用户注册绑定信息失败", err)
		return
	}

	// 检查用户是否已存在
	exist, err := dao.CheckUserExist(userInfo.Username)
	if err != nil {
		tools.StatusInternalErr(ctx, "检查用户是否存在发生错误", err)
		return
	}
	if exist {
		tools.StatusBadReq(ctx, "用户已存在", nil)
		return
	}

	if err = d.CreateUser(&userInfo); err != nil {
		tools.StatusInternalErr(ctx, "用户注册失败", err)
		return
	}

	tools.StatusSuccess(ctx, "创建用户成功", nil)
}

func UserLogin(ctx *gin.Context) {
	d := dao.NewUserRepository()
	var userReq types.UserLoginReq

	if ctx.ShouldBind(&userReq) == nil {
		// 如果绑定成功对form进行处理
		has, err := dao.CheckUserExist(userReq.Username)
		if err != nil {
			tools.StatusInternalErr(ctx, "用户登录绑定信息失败", err)
			return
		}
		if !has {
			tools.StatusBadReq(ctx, "用户未注册", nil)
			return
		}
		// 存在该用户，拿到该用户
		user, err := d.UserLogin(&userReq)
		if err != nil {
			tools.StatusInternalErr(ctx, "后端获取用户失败", err)
			return
		}

		// 验证密码是否一致
		if user.Password != utils.Md5(userReq.Password) {
			tools.StatusBadReq(ctx, "用户输入的密码错误", nil)
			return
		}
		// 2、生成token
		token, err := utils.GenerateToken(user.Identity, user.Name, global.TokenExpire)
		if err != nil {
			tools.StatusInternalErr(ctx, "生成用户token失败", nil)
			return
		}

		resp := new(types.UserLoginResp)
		resp.Token = token

		tools.StatusSuccess(ctx, "用户登录成功，返回token信息", resp)
		return
	}
}

func UserDetail(ctx *gin.Context) {
	d := dao.NewUserRepository()
	var userDetail types.UserDetailReq

	identity := ctx.Query("identity")
	fmt.Println("identity = ", identity)
	userDetail.Identity = identity
	userInfo, err := d.GetUserDetail(&userDetail)
	if err != nil {
		tools.StatusInternalErr(ctx, "获取用户数据失败", nil)
		return
	}
	resp := types.UserDetailResp{
		Username: userInfo.Name,
		Phone:    userInfo.Phone,
	}

	tools.StatusSuccess(ctx, "查询用户信息成功", resp)

	return
}

func DeleteUser(ctx *gin.Context) {

}
