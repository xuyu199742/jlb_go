package protocols

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 22222
	SUCCESS = 10000
)

func Result(code int, data interface{}, msg string, c *gin.Context, httpCode int) {
	// 开始时间
	c.JSON(httpCode, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context, httpCode int) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c, httpCode)
}

func OkWithMessage(message string, c *gin.Context, httpCode int) {
	Result(SUCCESS, map[string]interface{}{}, message, c, httpCode)
}

func OkWithData(data interface{}, c *gin.Context, httpCode int) {
	Result(SUCCESS, data, "操作成功", c, httpCode)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context, httpCode int) {
	Result(SUCCESS, data, message, c, httpCode)
}

func Fail(c *gin.Context, httpCode int) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c, httpCode)
}

func FailWithMessage(code int, message string, c *gin.Context, httpCode int) {
	Result(code, map[string]interface{}{}, message, c, httpCode)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context, httpCode int) {
	Result(ERROR, data, message, c, httpCode)
}
