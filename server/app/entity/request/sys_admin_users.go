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
