package toolkit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 99999
	FAILED  int = 10001
)

// 数据交互对象
type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func FastSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "成功",
	})
}

func FastFailed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Result{
		Code: FAILED,
		Msg:  "请求失败",
	})
}

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "成功",
		Data: v,
	})
}

func Failed(ctx *gin.Context, s string) {
	ctx.JSON(http.StatusOK, Result{
		Code: FAILED,
		Msg:  s,
	})
}

func FailedInfo(ctx *gin.Context, code int, s string) {
	ctx.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  s,
	})
}
