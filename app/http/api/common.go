package api

import (
	"encoding/json"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Status: http.StatusOK,
		Msg:    "Pong",
	})
}

// 把err转换为json
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: e.Error,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	} else {
		return serializer.Response{
			Status: e.Error,
			Msg:    "参数错误",
			Error:  err.Error(),
		}
	}
}
