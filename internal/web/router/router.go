package router

import (
	"file-system/internal/web/controller"
	"file-system/internal/web/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, types.HTTPResp{
				ErrCode: http.StatusOK,
				ErrMsg:  "system is healthy ~",
				Data:    nil,
			})
		})

		// 用户模块
		user := v1.Group("/user")
		{
			user.POST("/register", controller.UserRegister)
			user.POST("/login", controller.UserLogin)
			user.DELETE("/deregister", controller.DeleteUser)
		}

	}

	return router
}
