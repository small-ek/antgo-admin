package vo

type ResponseCaptcha struct {
	Id  string `json:"id" comment:"验证码标识"`  //验证码标识
	Pic string `json:"pic" comment:"验证码图片"` //验证码图片
}
