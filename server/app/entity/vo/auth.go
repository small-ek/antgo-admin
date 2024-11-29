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
