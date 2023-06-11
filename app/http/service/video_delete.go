package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/model"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
	"strconv"
)

type DeleteVideoService struct {
}

func (service *DeleteVideoService) Delete(ctx context.Context, uId uint, vId string) serializer.Response {
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
	err = videoDao.DeleteVideoByIDANDBossId(video.BossId, video.ID)
	if err != nil {
		code = e.ErrorVideoDelete
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "删除成功！",
	}

}
