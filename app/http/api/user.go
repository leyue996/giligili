package api

import (
	"giligili/app/http/pkg/e"
	"giligili/app/http/pkg/util"
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

func UserUpdatePw(c *gin.Context) {
	updateUserService := service.UpdateUserServicePw{}
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&updateUserService); err == nil {
		res := updateUserService.UserUpdatePw(c.Request.Context(), claims.Id, claims)
		c.JSON(e.Success, res)
		//重定向到登录
		//c.Redirect(http.StatusFound, "login")
	} else {
		c.JSON(e.Error, ErrorResponse(err))
	}
}
func UserUpdateNickName(c *gin.Context) {
	updateUserServiceNickName := service.UpdateUserServiceNickName{}
	claims, _ := util.ParseToken(c.GetHeader("access_token"))
	if err := c.ShouldBind(&updateUserServiceNickName); err == nil {
		res := updateUserServiceNickName.UserUpdateNickName(c.Request.Context(), claims.Id)
		c.JSON(e.Success, res)
	} else {
		c.JSON(e.Error, ErrorResponse(err))
	}
}
