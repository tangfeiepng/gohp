package response

import (
	"Walker/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

//响应正确请求

func Success(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, JsonResult{
		Code: 200,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format(global.Date),
	})
	ctx.Abort()
}

// 响应错误请求

func Error(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, JsonResult{
		Code: 400,
		Msg:  msg,
		Data: nil,
		Time: time.Now().Format(global.Date),
	})
	ctx.Abort()
}

func AuthError(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code": 403,
		"msg":  err,
		"data": "",
		"time": time.Now().Format(global.Date),
	})
	ctx.Abort()
}
func ExceedLimit(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code": 403,
		"msg":  "接口请求过于频繁",
		"data": "",
		"time": time.Now().Format(global.Date),
	})
	ctx.Abort()
}
