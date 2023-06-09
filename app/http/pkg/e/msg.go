package e

var MsgMap = map[int]string{
	Success: "ok",
	Error:   "错误",

	SuccessCreateUser: "创建用户成功",

	ErrorSelect: "查询数据库错误！",
	ErrorCreate: "创建失败！",

	ErrorRegisterPasswordRepeat:     "两次输入的密码不相同！",
	ErrorRegisterUserNameRepeat:     "用户名已注册！",
	ErrorRegisterPasswordEncryption: "密码加密失败！",
	ErrorLoginUserName:              "用户名不存在，请先注册！",
	ErrorLoginPassword:              "密码错误！",
}

func GetMsg(i int) string {
	str, ok := MsgMap[i]
	if !ok {
		return MsgMap[Error]
	}
	return str
}
