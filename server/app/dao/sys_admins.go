package dao

import (
	"github.com/small-ek/antgo/db/adb/sql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/models"
	
)

type SysAdminsDao struct {
	db    *gorm.DB
	models *models.SysAdmins
}

func NewSysAdminsDao() *SysAdminsDao {
	return &SysAdminsDao{db: ant.Db()}
}

// Create
func (dao *SysAdminsDao) Create(sys_admins *models.SysAdmins) error {
	return dao.db.Create(&sys_admins).Error
}

// DeleteById
func (dao *SysAdminsDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysAdminsDao) Update(sys_admins models.SysAdmins) error {
	return dao.db.Updates(&sys_admins).Error
}

// GetList
func (dao *SysAdminsDao) GetList() (list []models.SysAdmins) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysAdminsDao) GetPage(page page.PageParam, sys_admins models.SysAdmins) (list []models.SysAdmins, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		
		sql.Filters(page.Filter),
		sql.Order(page.Order, page.Desc),
		sql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(0).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysAdminsDao) GetById(id int) (row models.SysAdmins) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
