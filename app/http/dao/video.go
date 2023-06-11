package dao

import (
	"context"
	"giligili/app/http/model"
	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

func NewVideoDao(ctx context.Context) *VideoDao {
	return &VideoDao{NewDbClient(ctx)}
}
func (dao *VideoDao) CreateVideo(video *model.Video) error {
	return dao.DB.Model(&model.Video{}).Create(&video).Error
}

func (dao *VideoDao) ShowVideoById(vid uint) (video *model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("id=?", vid).First(&video).Error
	return
}
func (dao *VideoDao) ListVideo() (video []*model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Find(&video).Error
	return
}
func (dao *VideoDao) ShowVideoByIdANDBossId(uId, vid uint) (video *model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("boss_id=? AND id=?", uId, vid).First(&video).Error
	return
}
func (dao *VideoDao) UpdateVideoByIdANDBossId(video *model.Video, uId, vid uint) error {
	return dao.DB.Model(&model.Video{}).Where("boss_id=? AND id=?", uId, vid).Updates(&video).Error
}
func (dao *VideoDao) DeleteVideoByIDANDBossId(uId, vid uint) error {
	return dao.DB.Model(&model.Video{}).Where("boss_id=? AND id=?", uId, vid).Delete(&model.Video{}).Error
}
