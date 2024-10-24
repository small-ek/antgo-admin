package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/models"
)

type SysMenuRequest struct {
	models.SysMenu
	page.PageParam
}

type SysMenuRequestForm struct {
	models.SysMenu
	Title string `json:"title" form:"title" binding:"required" comment:"菜单名称"` //菜单名称

}
