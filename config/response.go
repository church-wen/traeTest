package config

import (
	"code.byted.org/zhuchaowen/trae/config/constants"
	"github.com/gin-gonic/gin"
)

// Response 统一返回结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应辅助函数
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    constants.ErrOK.Code,
		Message: constants.ErrOK.Message,
		Data:    data,
	})
}

// Fail 失败响应辅助函数
func Fail(c *gin.Context, errCode constants.ErrorCode, errMsg string) {
	if errMsg == "" {
		errMsg = errCode.Message
	}
	c.JSON(errCode.Code, Response{
		Code:    errCode.Code,
		Message: errMsg,
	})
}
