package api

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/vo"
	"server/app/entity/request"
	"server/app/service"
)

type SysAdminsController struct {
	vo.Base
	SysAdminsService *service.SysAdmins
}

func NewSysAdminsController() *SysAdminsController {
	return &SysAdminsController{
		SysAdminsService: service.NewSysAdminsService(),
	}
}

//	@Tags			管理员用户
//	@Summary		获取管理员用户分页数据
//	@Accept			json
//	@Produce		json
//	@Param		    data query request.SysAdminsRequest true "分页参数"
//	@Success		200	{object}	response.Write{data=response.Page{items=[]models.SysAdmins}}
//	@Failure		422	{object}	response.Write
//	@Router			/sys_admins [get] 路由
//
// Index 分页列表
func (ctrl *SysAdminsController) Index(c *gin.Context) {
	req := request.SysAdminsRequest{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	list, total, err := ctrl.SysAdminsService.SetReq(req).Index()
	if err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}
	ctrl.Success(c, "SUCCESS", ctrl.Page(total, list))
}

//	@Tags			管理员用户
//	@Summary		获取管理员用户详情数据
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Write{data=models.SysAdmins}
//	@Failure		422	{object}	response.Write
//	@Router			/sys_admins/:id [get]
//
// Show 详情
func (ctrl *SysAdminsController) Show(c *gin.Context) {
	var req request.SysAdminsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	row := ctrl.SysAdminsService.SetReq(req).Show()
	ctrl.Success(c, "SUCCESS", row)
}

//	@Tags			管理员用户
//	@Summary		创建管理员用户数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminsRequestForm true "创建参数"
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys_admins [post]
//
// Create 创建
func (ctrl *SysAdminsController) Create(c *gin.Context) {
	var req request.SysAdminsRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	if err := ctrl.SysAdminsService.SetReqForm(req).Store(); err != nil {
		ctrl.Fail(c, "CREATION_FAILED", err.Error())
		return
	}
	ctrl.Success(c, "SUCCESS")
}

//	@Tags			管理员用户
//	@Summary		修改管理员用户数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminsRequestForm true "更新参数"
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys_admins/:id [put]
//
// Update 修改
func (ctrl *SysAdminsController) Update(c *gin.Context) {
	var req request.SysAdminsRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	if err := ctrl.SysAdminsService.SetReqForm(req).Update();err!=nil{
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
//	@Router			/sys_admins/:id [delete]
//
// Delete 删除
func (ctrl *SysAdminsController) Delete(c *gin.Context) {
	var req request.SysAdminsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}

	if err := ctrl.SysAdminsService.SetReq(req).Delete(); err != nil {
		ctrl.Fail(c, "DELETE_FAILED", err.Error())
		return
	}

	ctrl.Success(c, "SUCCESS")
}

