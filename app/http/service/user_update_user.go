package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
)

type UpdateUserServiceNickName struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=30"`
}

func (service *UpdateUserServiceNickName) UserUpdateNickName(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.SelectUserByUserId(uId)

	user.NickName = service.NickName

	err = userDao.UpdateUserById(user, uId)
	if err != nil {
		code = e.ErrorUserUpdate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}

}
