package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/models"
)

type SysAdminUsersRequest struct {
	models.SysAdminUsers
	page.PageParam
}

type SysAdminUsersRequestForm struct {
	models.SysAdminUsers
	Username string `json:"username" form:"username" binding:"required" comment:"用户名"` //用户名
	Password string `json:"password" form:"password" binding:"required" comment:"密码"`  //密码
}

type SysAdminUsersLoginRequestForm struct {
	Username string `json:"username" form:"username" binding:"required" comment:"用户名"` //用户名
	Password string `json:"password" form:"password" binding:"required" comment:"密码"`  //密码
	Code     string `json:"code" form:"code" binding:"required" comment:"验证码"`         //验证码
	Id       string `json:"id" form:"id" binding:"required" comment:"验证码标识"`           //验证码标识
}
