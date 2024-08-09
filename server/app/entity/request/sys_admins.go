package request

import (
	"github.com/small-ek/antgo/utils/page"
	"server/app/models"
)

type SysAdminsRequest struct {
	models.SysAdmins
	page.PageParam
}

type SysAdminsRequestForm struct {
	models.SysAdmins
    
}
