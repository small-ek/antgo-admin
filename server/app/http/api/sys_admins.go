package api

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/request"
	"server/app/entity/vo"
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

//	@Summary		获取管理员用户分页数据
//	@Description	获取管理员用户分页
//	@Tags			SysAdmins
//	@Accept			json
//	@Produce		json
//	@Param		    data query request.SysAdminsRequest true "分页获取列表"
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/sys_admins [get] 路由
//
// Index 分页列表
func (ctrl *SysAdminsController) Index(c *gin.Context) {
	req := request.SysAdminsRequest{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	list, total, err := ctrl.SysAdminsService.SetReq(req).Index()
	if err != nil {
		ctrl.Fail(c, "异常错误", err.Error())
		return
	}
	ctrl.Success(c, "success", ctrl.Page(total, list))
}

//	@Summary		获取管理员用户单条数据
//	@Description	获取管理员用户详情
//	@Tags			SysAdmins
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/sys_admins/:id [get]
//
// Show 详情
func (ctrl *SysAdminsController) Show(c *gin.Context) {
	var req request.SysAdminsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	row := ctrl.SysAdminsService.SetReq(req).Show()
	ctrl.Success(c, "success", row)
}

//	@Summary		创建管理员用户数据
//	@Description	创建管理员用户数据
//	@Tags			SysAdmins
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminsRequestForm true "创建"
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/sys_admins [post]
//
// Create 创建
func (ctrl *SysAdminsController) Create(c *gin.Context) {
	var req request.SysAdminsRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	if err := ctrl.SysAdminsService.SetReqForm(req).Store(); err != nil {
		ctrl.Fail(c, "创建失败", err.Error())
		return
	}
	ctrl.Success(c, "success")
}

//	@Summary		修改管理员用户数据
//	@Description	修改管理员用户数据
//	@Tags			SysAdmins
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminsRequestForm true "更新"
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/sys_admins/:id [put]
//
// Update 修改
func (ctrl *SysAdminsController) Update(c *gin.Context) {
	var req request.SysAdminsRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	if err := ctrl.SysAdminsService.SetReqForm(req).Update(); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
	}
	ctrl.Success(c, "success")
}

//	@Summary		删除管理员用户数据
//	@Description	删除管理员用户数据
//	@Tags			SysAdmins
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/sys_admins/:id [delete]
//
// Delete 删除
func (ctrl *SysAdminsController) Delete(c *gin.Context) {
	var req request.SysAdminsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, "参数错误", err.Error())
		return
	}

	if err := ctrl.SysAdminsService.SetReq(req).Delete(); err != nil {
		ctrl.Fail(c, "删除失败", err.Error())
		return
	}

	ctrl.Success(c, "success")
}
