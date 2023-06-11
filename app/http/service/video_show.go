package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/model"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
	"strconv"
)

type ShowVideoService struct {
}

func (service *ShowVideoService) Show(ctx context.Context, vid string) serializer.Response {
	code := e.Success
	videoId, err := strconv.Atoi(vid)
	if err != nil {
		code = e.ErrorAtoI
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	videoDao := dao.NewVideoDao(ctx)
	video := new(model.Video)
	video, err = videoDao.ShowVideoById(uint(videoId))
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
		Data:   serializer.BuildVideo(video),
	}
}
