package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/model"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
	"strconv"
)

type UpdateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Info  string `json:"info" form:"info" binding:"required,max=300"`
}

func (service *UpdateVideoService) Update(ctx context.Context, uId uint, vId string) serializer.Response {
	code := e.Success
	videoId, err := strconv.Atoi(vId)
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
	video, err = videoDao.ShowVideoByIdANDBossId(uId, uint(videoId))
	if err != nil {
		code = e.ErrorVideoShow
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	err = videoDao.UpdateVideoByIdANDBossId(video, video.BossId, video.ID)
	if err != nil {
		code = e.ErrorVideoUpdate
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
