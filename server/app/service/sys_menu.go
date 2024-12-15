package service

import (
	"server/app/dao"
	"server/app/entity/models"
	"server/app/entity/request"
)

type SysMenu struct {
	req     request.SysMenuRequest
	reqForm request.SysMenuRequestForm
}

func NewSysMenuService() *SysMenu {
	return &SysMenu{}
}

// SetReq 设置参数
func (svc *SysMenu) SetReq(req request.SysMenuRequest) *SysMenu {
	svc.req = req
	return svc
}

// SetReqForm 设置参数
func (svc *SysMenu) SetReqForm(req request.SysMenuRequestForm) *SysMenu {
	req.SysMenu.Title = req.Title

	svc.reqForm = req
	return svc
}

// Index 分页
func (svc *SysMenu) Index() ([]models.SysMenu, int64, error) {
	return dao.NewSysMenuDao().GetPage(svc.req.PageParam)
}

// Show 查询单个
func (svc *SysMenu) Show() models.SysMenu {
	return dao.NewSysMenuDao().GetById(svc.req.SysMenu.Id)
}

// Store 添加
func (svc *SysMenu) Store() error {
	return dao.NewSysMenuDao().Create(&svc.reqForm.SysMenu)
}

// Update 修改
func (svc *SysMenu) Update() error {
	return dao.NewSysMenuDao().Update(svc.reqForm.SysMenu)
}

// Delete 删除
func (svc *SysMenu) Delete() error {
	return dao.NewSysMenuDao().DeleteById(svc.req.SysMenu.Id)
}
