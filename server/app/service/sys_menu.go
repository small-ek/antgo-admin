package service

import (
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/utils/conv"
	"go.uber.org/zap"
	"reflect"
	"server/app/dao"
	"server/app/entity/models"
	"server/app/entity/request"
)

type SysMenu struct {
	req     request.SysMenuRequest
	reqForm request.SysMenuRequestForm
	reqIds  request.IdsRequest
}

func NewSysMenuService() *SysMenu {
	return &SysMenu{}
}

// SetReq 设置参数
func (svc *SysMenu) SetReq(req interface{}) *SysMenu {
	switch value := req.(type) {
	case request.SysMenuRequest:
		svc.req = value
	case request.SysMenuRequestForm:
		conv.ToStruct(value, &svc.reqForm.SysMenu)
	case request.IdsRequest:
		svc.reqIds = value
	default:
		alog.Write.Error("SetReq", zap.Any("Unsupported request type", reflect.TypeOf(value)))
	}
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

// Deletes 批量删除
func (svc *SysMenu) Deletes() error {
	return dao.NewSysMenuDao().DeleteByIds(svc.reqIds.Ids)
}
