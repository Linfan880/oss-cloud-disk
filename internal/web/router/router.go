package router

import (
	"file-system/internal/web/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})

	router.POST("/api/v1/user", controller.CreateUser)

	return router
}
