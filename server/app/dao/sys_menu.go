package dao

import (
	"github.com/small-ek/antgo/db/adb/sql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type SysMenuDao struct {
	db     *gorm.DB
	models *models.SysMenu
}

func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{db: ant.Db()}
}

// Create
func (dao *SysMenuDao) Create(sysMenu *models.SysMenu) error {
	return dao.db.Create(&sysMenu).Error
}

// DeleteById
func (dao *SysMenuDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysMenuDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysMenuDao) Update(sysMenu models.SysMenu) error {
	return dao.db.Updates(&sysMenu).Error
}

// GetList
func (dao *SysMenuDao) GetList() (list []models.SysMenu) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysMenuDao) GetPage(page page.PageParam) (list []models.SysMenu, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		sql.Where("title", "LIKE", page.FilterMap["title"]),
		sql.Where("path", "LIKE", page.FilterMap["path"]),
		sql.Where("component", "LIKE", page.FilterMap["component"]),
		sql.Filters(page.Filter),
		sql.Order(page.Order, page.Desc),
		sql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(0).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysMenuDao) GetById(id int) (row models.SysMenu) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
