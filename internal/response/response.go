package response

import (
	"go-web-scaffold/internal/errors"

	"github.com/gin-gonic/gin"
)

// 定义响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 成功响应
func Success(data interface{}) *Response {
	return &Response{
		Code: errors.CodeSuccess,
		Msg:  errors.GetMsg(errors.CodeSuccess),
		Data: data,
	}
}

// 错误响应
func Error(code int, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  errors.GetMsg(code),
		Data: data,
	}
}

// 统一 JSON 响应
func JSONResponse(c *gin.Context, httpCode int, response *Response) {
	c.JSON(httpCode, response)
}
