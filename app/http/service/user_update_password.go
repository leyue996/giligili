package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/pkg/e"
	"giligili/app/http/pkg/util"
	"giligili/app/http/serializer"
)

type UpdateUserServicePw struct {
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`

	//新输入密码
	PasswordNew        string `form:"password_new" json:"password_new" binding:"required,min=8,max=40"`
	PasswordNewConfirm string `form:"password_new_confirm" json:"password_new_confirm" binding:"required,min=8,max=40"`
}

func (service *UpdateUserServicePw) UserUpdatePw(ctx context.Context, uId uint, claims *util.Claims) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.SelectUserByUserId(uId)

	if ok := user.CheckPassword(service.Password); !ok {
		code = e.ErrorLoginPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if service.PasswordNew != service.PasswordNewConfirm {
		code = e.ErrorNewPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = user.SetPassword(service.PasswordNew)
	if err != nil {
		code = e.ErrorRegisterPasswordEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = userDao.UpdateUserById(user, uId)
	if err != nil {
		code = e.ErrorUserUpdate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//旧token设为过期
	claims.ExpireToken()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "重定向login",
	}
}
