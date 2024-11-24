package service

import (
	"errors"
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/jwt"
	"go.uber.org/zap"
	"reflect"
	"server/app/dao"
	"server/app/entity/models"
	"server/app/entity/request"
	"server/utils"
	"time"
)

type SysAdminUsers struct {
	req                  request.SysAdminUsersRequest
	reqForm              request.SysAdminUsersRequestForm
	reqSysAdminUsersInfo request.SysAdminUsersInfoRequest
	reqLoginForm         request.SysAdminUsersLoginRequestForm
}

func NewSysAdminUsersService() *SysAdminUsers {
	return &SysAdminUsers{}
}

// SetReq 设置参数，支持不同类型的请求
func (svc *SysAdminUsers) SetReq(req interface{}) *SysAdminUsers {
	switch value := req.(type) {
	case request.SysAdminUsersRequest:
		svc.req = value
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
	default:
		// 不支持的类型，抛出异常或者返回错误
		alog.Write.Error("SetReq", zap.Any("Unsupported request type", reflect.TypeOf(value)))
	}
	return svc
}

// Index 分页
func (svc *SysAdminUsers) Index() ([]models.SysAdminUsers, int64, error) {
	return dao.NewSysAdminUsersDao().GetPage(svc.req.PageParam, svc.req.SysAdminUsers)
}

// Show 查询单个
func (svc *SysAdminUsers) Show() models.SysAdminUsers {
	return dao.NewSysAdminUsersDao().GetById(svc.req.SysAdminUsers.Id)
}

// Store 添加
func (svc *SysAdminUsers) Store() error {
	return dao.NewSysAdminUsersDao().Create(&svc.reqForm.SysAdminUsers)
}

// Update 修改
func (svc *SysAdminUsers) Update() error {
	return dao.NewSysAdminUsersDao().Update(svc.reqForm.SysAdminUsers)
}

// Delete 删除
func (svc *SysAdminUsers) Delete() error {
	return dao.NewSysAdminUsersDao().DeleteById(svc.req.SysAdminUsers.Id)
}

// Login 添加
func (svc *SysAdminUsers) Login() (result map[string]interface{}, err error) {
	item := dao.NewSysAdminUsersDao().GetByUserName(svc.reqLoginForm.Username)
	if utils.VerifyPassword(item.Password, svc.reqLoginForm.Password) == true {

		token, expiresAt, err := svc.token(map[string]interface{}{"id": item.Id, "username": item.Username})
		if err != nil {
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
	return nil, errors.New("LOGIN_ERROR")
}

// token 生成token
func (svc *SysAdminUsers) token(str map[string]interface{}) (token string, expiresAt int64, err error) {
	var j = jwt.New().SetPrivateKey(conv.Bytes(config.GetString("jwt.private_key")))
	expiresAt = time.Now().Add(time.Hour * time.Duration(config.GetInt64("jwt.exp"))).Unix()
	token, err = j.Encrypt(str, expiresAt)
	return
}
