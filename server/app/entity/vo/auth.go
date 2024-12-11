package vo

// 获取验证码失败信息提示，返回错误信息
func GetUpdatePasswordErrorMsg(err error) string {
	switch err.Error() {
	case PASSWORD_ERROR:
		return PASSWORD_ERROR
	default:
		return UPDATE_FAILED
	}
}

// CreateUserErrorMsg 创建用户失败信息提示，返回错误信息
func CreateUserErrorMsg(err error) string {
	switch err.Error() {
	case USERNAME_EXISTS:
		return USERNAME_EXISTS
	default:
		return CREATION_FAILED
	}
}

// UpdateUserErrorMsg 修改用户失败信息提示，返回错误信息
func UpdateUserErrorMsg(err error) string {
	switch err.Error() {
	case USERNAME_EXISTS:
		return USERNAME_EXISTS
	default:
		return UPDATE_FAILED
	}
}
