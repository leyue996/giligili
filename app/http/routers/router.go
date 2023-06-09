package routers

import (
	"giligili/app/http/api"
	"giligili/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors()) //加载中间件

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", api.Ping)
	}
	return r
}
