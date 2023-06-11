package e

const (
	Success           = 200
	SuccessCreateUser = 20001

	Error = 400

	ErrorSelect = 41000
	ErrorCreate = 41001
	ErrorAtoI   = 41002

	ErrorAuthToken             = 42000
	ErrorAuthCheckTokenTimeout = 42001

	ErrorRegisterPasswordRepeat     = 40000
	ErrorRegisterUserNameRepeat     = 40001
	ErrorRegisterPasswordEncryption = 40002
	ErrorLoginUserName              = 40003
	ErrorLoginPassword              = 40004
	ErrorNewPassword                = 40005
	ErrorUserUpdate                 = 40006

	ErrorVideoShow   = 43000
	ErrorVideoUpdate = 43001
	ErrorVideoDelete = 43002
)
