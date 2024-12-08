package service

import (
	"errors"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/jwt"
	"go.uber.org/zap"
	"reflect"
	"server/app/dao"
	"server/app/entity/models"
	"server/app/entity/request"
	"server/app/entity/vo"
	"server/utils"
	"time"
)

type SysAdminUsers struct {
	req                  request.SysAdminUsersRequest
	reqForm              request.SysAdminUsersRequestForm
	reqSysAdminUsersInfo request.SysAdminUsersInfoRequest
	reqLoginForm         request.SysAdminUsersLoginRequestForm
	reqPasswordForm      request.SysAdminUsersPasswordRequestForm
	reqIds               request.IdsRequest
}

func NewSysAdminUsersService() *SysAdminUsers {
	return &SysAdminUsers{}
}

// SetReq 设置参数，支持不同类型的请求
func (svc *SysAdminUsers) SetReq(req interface{}) *SysAdminUsers {
	switch value := req.(type) {
	case request.SysAdminUsersRequest:
		svc.req = value
	case request.IdsRequest:
		svc.reqIds = value
	case request.SysAdminUsersRequestForm:
		svc.reqForm = value
		// 根据需求做字段赋值
		svc.reqForm.SysAdminUsers.Username = value.Username
		svc.reqForm.SysAdminUsers.Password = value.Password
		svc.reqForm.SysAdminUsers.NickName = value.NickName
	case request.SysAdminUsersInfoRequest:
		conv.ToStruct(value, &svc.reqForm.SysAdminUsers)
	case request.SysAdminUsersLoginRequestForm:
		svc.reqLoginForm = value

	case request.SysAdminUsersPasswordRequestForm:
		svc.reqForm.SysAdminUsers.Password = value.Password
		svc.reqForm.SysAdminUsers.Id = value.Id
		svc.reqPasswordForm = value
	default:
		alog.Write.Error("SetReq", zap.Any("Unsupported request type", reflect.TypeOf(value)))
	}
	return svc
}

// Index 分页
func (svc *SysAdminUsers) Index() ([]models.SysAdminUsers, int64, error) {
	return dao.NewSysAdminUsersDao().GetPage(svc.req.PageParam)
}

// Show 查询单个
func (svc *SysAdminUsers) Show() models.SysAdminUsers {
	return dao.NewSysAdminUsersDao().GetById(svc.req.SysAdminUsers.Id)
}

// Store 添加
func (svc *SysAdminUsers) Store() error {
	return dao.NewSysAdminUsersDao().Create(&svc.reqForm.SysAdminUsers)
}

// UpdatePassword 修改密码
func (svc *SysAdminUsers) UpdatePassword() error {
	item := dao.NewSysAdminUsersDao().GetById(svc.reqForm.SysAdminUsers.Id)

	if utils.VerifyPassword(item.Password, svc.reqForm.SysAdminUsers.Password) == true {
		password, err := utils.GeneratePassword(svc.reqPasswordForm.NewPassword)
		if err != nil {
			return err
		}
		svc.reqForm.SysAdminUsers.Password = password
		return dao.NewSysAdminUsersDao().Update(svc.reqForm.SysAdminUsers)
	}
	return errors.New(vo.PASSWORD_ERROR)
}

// Update 修改
func (svc *SysAdminUsers) Update() error {
	return dao.NewSysAdminUsersDao().Update(svc.reqForm.SysAdminUsers)
}

// Deletes 批量删除
func (svc *SysAdminUsers) Deletes() error {
	return dao.NewSysAdminUsersDao().DeleteByIds(svc.reqIds.Ids)
}

// Login 登录操作
func (svc *SysAdminUsers) Login() (result map[string]interface{}, err error) {
	item := dao.NewSysAdminUsersDao().GetByUserName(svc.reqLoginForm.Username)
	if utils.VerifyPassword(item.Password, svc.reqLoginForm.Password) == true {

		token, expiresAt, err := svc.token(map[string]interface{}{"id": item.Id, "username": item.Username, "device-id": svc.reqLoginForm.DeviceId})
		if err != nil {
			return nil, err
		}

		if err = ant.Redis().Set("sys_admin_users:"+conv.String(item.Id), token, int64(3600*1000*time.Duration(config.GetInt64("jwt.exp")))); err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"token":     token,
			"expiresAt": expiresAt,
			"user": map[string]interface{}{
				"username":  item.Username,
				"nick_name": item.NickName,
				"phone":     item.Phone,
				"email":     item.Email,
				"status":    item.Status,
			}}, nil
	}
	return nil, errors.New(vo.LOGIN_ERROR)
}

// token 生成token
func (svc *SysAdminUsers) token(str map[string]interface{}) (token string, expiresAt int64, err error) {
	var j = jwt.New().SetPrivateKey(conv.Bytes(config.GetString("jwt.private_key")))
	expiresAt = time.Now().Add(time.Hour * time.Duration(config.GetInt64("jwt.exp"))).Unix()
	token, err = j.Encrypt(str, expiresAt)
	return
}
