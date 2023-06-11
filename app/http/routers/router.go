package routers

import (
	"giligili/app/http/api"
	"giligili/app/http/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())                    //加载中间件
	r.StaticFS("/static", http.Dir("./static")) //加载静态文件

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", api.Ping)

		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		v1.GET("video/:id", api.ShowVideo) //具体视频
		v1.GET("videos", api.ListVideo)    //视频列表
		//登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.PUT("user/updatePW", api.UserUpdatePw)
			authed.PUT("user/updateNickName", api.UserUpdateNickName)

			authed.POST("video", api.CreateVideo)
			authed.PUT("video/:id", api.UpdateVideo)
			authed.DELETE("video/:id", api.DeleteVideo)
		}
	}
	return r
}
