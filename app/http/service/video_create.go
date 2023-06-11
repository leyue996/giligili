package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/model"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
)

type CreateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Info  string `json:"info" form:"info" binding:"required,max=300"`
}

func (service *CreateVideoService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.SelectUserByUserId(uId)
	if err != nil {
		code = e.ErrorSelect
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	video := new(model.Video)
	video.Title = service.Title
	video.Info = service.Info
	video.BossName = user.UserName
	video.BossId = user.ID

	videoDao := dao.NewVideoDao(ctx)
	err = videoDao.CreateVideo(video)
	if err != nil {
		code = e.ErrorCreate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildVideo(video),
	}
}
