package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/model"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
)

type ListVideosService struct {
}

func (service *ListVideosService) List(ctx context.Context) serializer.Response {
	code := e.Success
	videoDao := dao.NewVideoDao(ctx)
	var video []*model.Video
	video, err := videoDao.ListVideo()
	if err != nil {
		code = e.ErrorVideoShow
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildVideos(video),
	}
}
