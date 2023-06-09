package api

import (
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
