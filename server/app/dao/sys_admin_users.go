package dao

import (
	"fmt"
	"github.com/small-ek/antgo/db/adb/sql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type SysAdminUsersDao struct {
	db     *gorm.DB
	models *models.SysAdminUsers
}

func NewSysAdminUsersDao(db *gorm.DB) *SysAdminUsersDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysAdminUsersDao{db: db}
}

// Create
func (dao *SysAdminUsersDao) Create(sysAdminUsers *models.SysAdminUsers) error {
	return dao.db.Create(&sysAdminUsers).Error
}

// DeleteById
func (dao *SysAdminUsersDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysAdminUsersDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysAdminUsersDao) Update(sysAdminUsers models.SysAdminUsers) error {
	return dao.db.Updates(&sysAdminUsers).Error
}

// GetList
func (dao *SysAdminUsersDao) GetList() (list []models.SysAdminUsers) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysAdminUsersDao) GetPage(page page.PageParam) (list []models.SysAdminUsers, total int64, err error) {
	fmt.Println(conv.String(page.FilterMap))
	err = dao.db.Model(&dao.models).Scopes(
		sql.Where("username", "LIKE", page.FilterMap["username"]),
		sql.Where("nick_name", "LIKE", page.FilterMap["nick_name"]),
		sql.Where("phone", "LIKE", page.FilterMap["phone"]),
		sql.Where("email", "LIKE", page.FilterMap["email"]),
		sql.Where("status", "=", page.FilterMap["status"]),
		sql.Where("created_at", "BETWEEN", page.FilterMap["created_at"]),
		sql.Filters(page.Filter),
		sql.Order(page.Order, page.Desc),
		sql.Paginate(page.PageSize, page.CurrentPage),
	).Omit("password").Find(&list).Offset(0).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysAdminUsersDao) GetById(id int) (row *models.SysAdminUsers) {
	dao.db.Model(&dao.models).Omit("password").Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetByUserNameAndStatus
func (dao *SysAdminUsersDao) GetByUserNameAndStatus(username string) (row *models.SysAdminUsers) {
	dao.db.Model(&dao.models).Where("username=? AND status=2", username).Limit(1).Find(&row)
	return row
}

// GetByUserNameAndStatus
func (dao *SysAdminUsersDao) GetByUserName(username string) (row *models.SysAdminUsers) {
	dao.db.Model(&dao.models).Where("username=?", username).Limit(1).Find(&row)
	return row
}
