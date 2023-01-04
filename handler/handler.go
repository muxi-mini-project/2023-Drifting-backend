package handler

import (
	"github.com/gin-gonic/gin"
)

// Response 请求响应
type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
} //@name Response

func SendGoodResponse(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func SendBadResponse(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(400, Response{
		Code:    400,
		Message: message,
		Data:    data,
	})
}

func SendError(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(500, Response{
		Code:    500,
		Message: message,
		Data:    data,
	})
}
