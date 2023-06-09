package serializer

import (
	"giligili/app/http/model"
	"log"
)

type User struct {
	//给前端展示的数据
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`

	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *User {
	log.Println(user.PasswordDigest)
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,

		Status:   user.Status,
		Avatar:   user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
