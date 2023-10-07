package tools

import (
	"file-system/internal/web/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, types.HTTPResp{
		ErrCode: http.StatusOK,
		ErrMsg:  message,
		Data:    data,
	})
}

func StatusSimpleSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, types.HTTPResp{
		ErrCode: http.StatusOK,
		ErrMsg:  "success",
		Data:    nil,
	})
}
func StatusNotFound(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusNotFound, types.HTTPResp{
		ErrCode: http.StatusNotFound,
		ErrMsg:  message,
		Data:    data,
	})
}

func StatusInternalErr(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusInternalServerError, types.HTTPResp{
		ErrCode: http.StatusInternalServerError,
		ErrMsg:  message,
		Data:    data,
	})
}

func StatusBadReq(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusBadRequest, types.HTTPResp{
		ErrCode: http.StatusBadRequest,
		ErrMsg:  message,
		Data:    data,
	})
}
