package res

import (
	"github.com/gin-gonic/gin"
	"gvb_server/utils"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, context *gin.Context) {
	context.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, context *gin.Context) {
	Result(Success, data, msg, context)
}
func OkWithData(data any, context *gin.Context) {
	Result(Success, data, "success", context)
}
func OkWithList(list any, count int64, context *gin.Context) {
	OkWithData(ListResponse[any]{
		Count: count,
		List:  list,
	}, context)
}
func OkWithMessage(msg string, context *gin.Context) {
	Result(Success, map[string]any{}, msg, context)
}
func OkWithout(context *gin.Context) {
	Result(Success, map[string]any{}, "成功", context)
}

func Fail(data any, msg string, context *gin.Context) {
	Result(Error, data, msg, context)
}
func FailWithMessage(msg string, context *gin.Context) {
	Result(Error, map[string]any{}, msg, context)
}
func FailWithError(err error, obj any, context *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, context)
}
func FailWithCode(code ErrorCode, context *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, context)
		return
	}
	Result(Error, map[string]any{}, "Unknown Error", context)

}
