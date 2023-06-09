package dao

import (
	"context"
	"giligili/app/http/model"
	"gorm.io/gorm"
)

// 持久化db
type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDbClient(ctx)}
}

func (dao *UserDao) ExistsOrNotByUserName(username string) (count int64, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) SelectUserByUserName(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	return
}
