package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysAdminUsersRequest struct {
	models.SysAdminUsers
	page.PageParam
}

type SysAdminUsersRequestForm struct {
	models.SysAdminUsers
	Username string `json:"username" form:"username" binding:"required" comment:"用户名"` //用户名
}

type SysAdminUsersLoginRequestForm struct {
	Username string `json:"username" form:"username" binding:"required" comment:"用户名"` //用户名
	Password string `json:"password" form:"password" binding:"required" comment:"密码"`  //密码
	Code     string `json:"code" form:"code" binding:"required" comment:"验证码"`         //验证码
	DeviceId string `json:"device-id" form:"device-id" comment:"设备号"`                  //验证码
	Id       string `json:"id" form:"id" binding:"required" comment:"验证码标识"`           //验证码标识
}

type SysAdminUsersInfoRequest struct {
	NickName string `binding:"required" json:"nick_name" form:"nick_name" comment:"昵称"` //昵称
	Phone    string `json:"phone" form:"phone" comment:"手机"`                            //手机
	Email    string `json:"email" form:"email" comment:"电子邮件"`                          //电子邮件
	Id       int    `json:"id" form:"id" comment:"标识"`                                  //标识
}

type SysAdminUsersPasswordRequestForm struct {
	Id          int    `json:"id" form:"id" comment:"标识"`                                         //标识
	Password    string `binding:"required" json:"password" form:"password" comment:"密码"`          //密码
	NewPassword string `binding:"required" json:"new_password" form:"new_password" comment:"新密码"` //密码
}
