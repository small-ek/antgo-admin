package service

import (
	"server/app/dao"
	"server/app/models"
	"server/app/entity/request"
)

type SysAdmins struct {
	req request.SysAdminsRequest
	reqForm request.SysAdminsRequestForm
}

func NewSysAdminsService() *SysAdmins {
	return &SysAdmins{}
}

//SetReq 设置参数
func (svc *SysAdmins) SetReq(req request.SysAdminsRequest) *SysAdmins {
	svc.req = req
	return svc
}

// SetReqForm 设置参数
func (svc *SysAdmins) SetReqForm(req request.SysAdminsRequestForm) *SysAdmins {
	
	svc.reqForm = req
	return svc
}

// Index 分页
func (svc *SysAdmins) Index() ([]models.SysAdmins, int64, error) {
	return dao.NewSysAdminsDao().GetPage(svc.req.PageParam, svc.req.SysAdmins)
}

// Show 查询单个
func (svc *SysAdmins) Show() models.SysAdmins {
	return dao.NewSysAdminsDao().GetById(svc.req.SysAdmins.Id)
}

// Store 添加
func (svc *SysAdmins) Store() error {
	return dao.NewSysAdminsDao().Create(&svc.reqForm.SysAdmins)
}

// Update 修改
func (svc *SysAdmins) Update() error {
	return dao.NewSysAdminsDao().Update(svc.reqForm.SysAdmins)
}

// Delete 删除
func (svc *SysAdmins) Delete() error {
	return dao.NewSysAdminsDao().DeleteById(svc.req.SysAdmins.Id)
}

