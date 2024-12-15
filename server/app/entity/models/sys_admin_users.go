package models

import (
	"gorm.io/gorm"
	"server/utils"
	"time"
)

type SysAdminUsers struct {
	Id        int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"` //标识
	Username  string         `gorm:"column:username" json:"username" form:"username" comment:"用户名"`                 //用户名
	Password  string         `gorm:"column:password" json:"password" form:"password" comment:"密码"`                  //密码
	NickName  string         `gorm:"column:nick_name" json:"nick_name" form:"nick_name" comment:"昵称"`               //昵称
	Phone     string         `gorm:"column:phone" json:"phone" form:"phone" comment:"手机"`                           //手机
	Email     string         `gorm:"column:email" json:"email" form:"email" comment:"电子邮件"`                         //电子邮件
	Status    int            `gorm:"column:status" json:"status" form:"status" comment:"状态:0=禁用;1=启用"`              //状态:1=成功;2=失败
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`          //创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"-" form:"updated_at" comment:"修改时间"`                   //修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                   //删除时间
}

func (SysAdminUsers) TableName() string {
	return "antgo.sys_admin_users"
}

func (m *SysAdminUsers) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		pwd, err2 := utils.GeneratePassword(m.Password)
		if err2 != nil {
			return err2
		}
		m.Password = pwd
	}
	return nil
}

func (m *SysAdminUsers) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		pwd, err2 := utils.GeneratePassword(m.Password)
		if err2 != nil {
			return err2
		}
		m.Password = pwd
	}
	return nil
}
