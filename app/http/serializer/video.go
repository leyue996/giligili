package serializer

import "giligili/app/http/model"

type Video struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	Info     string `json:"info"`
	BossId   uint   `json:"boss_id"`
	BossName string `json:"boss_name"`
	CreateAt int64  `json:"create_at"`
}

func BuildVideo(item *model.Video) Video {
	return Video{
		Id:       item.ID,
		Title:    item.Title,
		Info:     item.Info,
		BossId:   item.BossId,
		BossName: item.BossName,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildVideos(items []*model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
