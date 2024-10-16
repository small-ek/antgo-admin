package api

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/utils/page"
	"server/app/entity/vo"
	"server/app/entity/request"
	"server/app/service"
)

type SysMenuController struct {
	vo.Base
	SysMenuService *service.SysMenu
}

func NewSysMenuController() *SysMenuController {
	return &SysMenuController{
		SysMenuService: service.NewSysMenuService(),
	}
}

//	@Tags			后台菜单
//	@Summary		获取后台菜单分页数据
//	@Accept			json
//	@Produce		json
//	@Param		    data query request.SysMenuRequest true "分页参数"
//	@Success		200	{object}	response.Write{data=response.Page{items=[]models.SysMenu}}
//	@Failure		422	{object}	response.Write
//	@Router			/sys-menu [get] 路由
//
// Index 分页列表
func (ctrl *SysMenuController) Index(c *gin.Context) {
	req := request.SysMenuRequest{
		PageParam: page.New(),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err.Error())
		return
	}

	list, total, err := ctrl.SysMenuService.SetReq(req).Index()
	if err != nil {
		ctrl.Fail(c, vo.FAILED, err.Error())
		return
	}
	ctrl.Success(c, vo.SUCCESS, ctrl.Page(total, list))
}

//	@Tags			后台菜单
//	@Summary		获取后台菜单详情数据
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Write{data=models.SysMenu}
//	@Failure		422	{object}	response.Write
//	@Router			/sys-menu/:id [get]
//
// Show 详情
func (ctrl *SysMenuController) Show(c *gin.Context) {
	var req request.SysMenuRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err.Error())
		return
	}

	result := ctrl.SysMenuService.SetReq(req).Show()
	ctrl.Success(c, vo.SUCCESS, result)
}

//	@Tags			后台菜单
//	@Summary		创建后台菜单数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysMenuRequestForm true "创建参数"
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys-menu [post]
//
// Create 创建
func (ctrl *SysMenuController) Create(c *gin.Context) {
	var req request.SysMenuRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err.Error())
		return
	}

	if err := ctrl.SysMenuService.SetReqForm(req).Store(); err != nil {
		ctrl.Fail(c, vo.CREATION_FAILED, err.Error())
		return
	}
	ctrl.Success(c, vo.CREATION_SUCCESS)
}

//	@Tags			后台菜单
//	@Summary		修改后台菜单数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysMenuRequestForm true "更新参数"
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys-menu/:id [put]
//
// Update 修改
func (ctrl *SysMenuController) Update(c *gin.Context) {
	var req request.SysMenuRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err.Error())
		return
	}

	if err := ctrl.SysMenuService.SetReqForm(req).Update();err!=nil{
		ctrl.Fail(c, vo.UPDATE_FAILED, err.Error())
	}
	ctrl.Success(c, vo.UPDATE_SUCCESS)
}

//	@Tags			后台菜单
//	@Summary		删除后台菜单数据
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/sys-menu/:id [delete]
//
// Delete 删除
func (ctrl *SysMenuController) Delete(c *gin.Context) {
	var req request.SysMenuRequest
	if err := c.ShouldBindUri(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err.Error())
		return
	}

	if err := ctrl.SysMenuService.SetReq(req).Delete(); err != nil {
		ctrl.Fail(c, vo.DELETE_FAILED, err.Error())
		return
	}

	ctrl.Success(c, vo.DELETE_SUCCESS)
}

