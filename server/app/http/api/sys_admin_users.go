package api

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/request"
	"server/app/entity/vo"
	"server/app/service"
)

type SysAdminUsersController struct {
	vo.Base
	SysAdminUsersService *service.SysAdminUsers
}

func NewSysAdminUsersController() *SysAdminUsersController {
	return &SysAdminUsersController{
		SysAdminUsersService: service.NewSysAdminUsersService(),
	}
}

//	@Tags			管理员用户
//	@Summary		获取管理员用户分页数据
//	@Accept			json
//	@Produce		json
//	@Param		    data query request.SysAdminUsersRequest true "分页参数"
//	@Success		200	{object}	response.Write{data=response.Page{items=[]models.SysAdminUsers}}
//	@Failure		422	{object}	response.Write
//	@Router			/sys-admin-users [get] 路由
//
// Index 分页列表
func (ctrl *SysAdminUsersController) Index(c *gin.Context) {
	req := request.SysAdminUsersRequest{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	list, total, err := ctrl.SysAdminUsersService.SetReq(req).Index()
	if err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}
	ctrl.SetCode(102).Success(c, "SUCCESS", ctrl.Page(total, list))
}

//	@Tags			管理员用户
//	@Summary		获取管理员用户详情数据
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Write{data=models.SysAdminUsers}
//	@Failure		422	{object}	response.Write
//	@Router			/sys-admin-users/:id [get]
//
// Show 详情
func (ctrl *SysAdminUsersController) Show(c *gin.Context) {
	var req request.SysAdminUsersRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	result := ctrl.SysAdminUsersService.SetReq(req).Show()
	ctrl.Success(c, "SUCCESS", result)
}

//	@Tags			管理员用户
//	@Summary		创建管理员用户数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminUsersRequestForm true "创建参数"
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys-admin-users [post]
//
// Create 创建
func (ctrl *SysAdminUsersController) Create(c *gin.Context) {
	var req request.SysAdminUsersRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	if err := ctrl.SysAdminUsersService.SetReqForm(req).Store(); err != nil {
		ctrl.Fail(c, "CREATION_FAILED", err.Error())
		return
	}
	ctrl.Success(c, "SUCCESS")
}

//	@Tags			管理员用户
//	@Summary		修改管理员用户数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminUsersRequestForm true "更新参数"
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys-admin-users/:id [put]
//
// Update 修改
func (ctrl *SysAdminUsersController) Update(c *gin.Context) {
	var req request.SysAdminUsersRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	if err := ctrl.SysAdminUsersService.SetReqForm(req).Update(); err != nil {
		ctrl.Fail(c, "UPDATE_FAILED", err.Error())
	}
	ctrl.Success(c, "SUCCESS")
}

//	@Tags			管理员用户
//	@Summary		删除管理员用户数据
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys-admin-users/:id [delete]
//
// Delete 删除
func (ctrl *SysAdminUsersController) Delete(c *gin.Context) {
	var req request.SysAdminUsersRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	if err := ctrl.SysAdminUsersService.SetReq(req).Delete(); err != nil {
		ctrl.Fail(c, "DELETE_FAILED", err.Error())
		return
	}

	ctrl.Success(c, "SUCCESS")
}
