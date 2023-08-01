package commonResult

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IResponse
// 声明返回函数
type IResponse interface {
	Success(context *gin.Context, msg string, data interface{})
	Error(context *gin.Context, msg string, data interface{})
}

// Result
// 多态返回实现类
func Result(Ir IResponse) IResponse {
	return Ir
}

// Response
// 简单返回
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (response *Response) Success(context *gin.Context, msg string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
	})
}

func (response *Response) Error(context *gin.Context, msg string, data interface{}) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"msg":  msg,
		"data": data,
	})
}

// CustomResponse
// 自定义返回
type CustomResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Ok  = response(http.StatusOK, "ok")
	Err = response(http.StatusInternalServerError, "error")
)

func response(code int, msg string) *CustomResponse {
	return &CustomResponse{
		code,
		msg,
		nil,
	}
}

func (custom *CustomResponse) Success(context *gin.Context, msg string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": custom.Code,
		"msg":  msg,
		"data": data,
	})
}

func (custom *CustomResponse) Error(context *gin.Context, msg string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": custom.Code,
		"msg":  msg,
		"data": data,
	})
}
