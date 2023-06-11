package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	NickName       string `gorm:"column:nick_name"`
	Status         string
	Avatar         string `gorm:"size:1000"`
}

const (
	PasswordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

// 加密密码
func (user *User) SetPassword(pw string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), PasswordCost)
	if err != nil {
		return err
	}

	user.PasswordDigest = string(bytes)
	return err

}

// 验证密码是否一致
func (user *User) CheckPassword(pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(pw))
	return err == nil
}
