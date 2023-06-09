package service

import (
	"context"
	"giligili/app/http/dao"
	"giligili/app/http/model"
	"giligili/app/http/pkg/e"
	"giligili/app/http/serializer"
)

type UserRegisterService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	//重复输入密码
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// 验证注册表单是否符合规范
func (service *UserRegisterService) ValidateForm(userDao *dao.UserDao) int {
	if service.PasswordConfirm != service.Password {
		return e.ErrorRegisterPasswordRepeat
	}
	count, err := userDao.ExistsOrNotByUserName(service.UserName)
	if err != nil {
		return e.ErrorSelect
	}
	if count != 0 {
		return e.ErrorRegisterUserNameRepeat
	}
	return e.Success
}

// 注册
func (service *UserRegisterService) Register(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	code := service.ValidateForm(userDao)

	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	var user model.User
	err := user.SetPassword(service.Password)
	if err != nil {
		code = e.ErrorRegisterPasswordEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.UserName = service.UserName
	user.Nickname = service.NickName
	user.Status = model.Active

	//log.Println(user)
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.ErrorCreate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	code = e.SuccessCreateUser
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
