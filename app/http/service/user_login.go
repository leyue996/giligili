package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/pkg/e"
	"giligili/app/http/pkg/util"
	"giligili/app/http/serializer"
)

type UserLoginService struct {
	UserName string ` form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *UserLoginService) Login(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	code := e.Success
	count, err := userDao.ExistsOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ErrorSelect
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if count == 0 {
		code = e.ErrorRegisterUserNameRepeat
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user, err := userDao.SelectUserByUserName(service.UserName)
	if ok := user.CheckPassword(service.Password); !ok {
		code = e.ErrorLoginPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//生产token
	token, err := util.GenerateToken(user.ID, user.UserName)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}
}
