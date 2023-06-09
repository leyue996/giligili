package api

import (
	"giligili/app/http/pkg/e"
	"giligili/app/http/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	userRegisterService := service.UserRegisterService{}
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(e.Success, res)
	} else {
		c.JSON(e.Error, ErrorResponse(err))
	}

}
func UserLogin(c *gin.Context) {
	userLoginService := service.UserLoginService{}
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(e.Success, res)
	} else {
		c.JSON(e.Error, ErrorResponse(err))
	}

}
