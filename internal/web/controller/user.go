package controller

import (
	"file-system/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	d := repository.NewUserRepository()
	if err := d.CreateUser("linfan", "lin2738963", "19724257473"); err != nil {
		ctx.JSON(500, "创建用户失败")
		return
	}
	ctx.JSON(200, "创建用户成功")
}
