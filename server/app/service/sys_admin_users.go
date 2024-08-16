package service

import (
	"errors"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/jwt"
	"server/app/dao"
	"server/app/entity/models"
	"server/app/entity/request"
	"server/utils"
	"time"
)

type SysAdminUsers struct {
	req          request.SysAdminUsersRequest
	reqForm      request.SysAdminUsersRequestForm
	reqLoginForm request.SysAdminUsersLoginRequestForm
}

func NewSysAdminUsersService() *SysAdminUsers {
	return &SysAdminUsers{}
}

// SetReq 设置参数
func (svc *SysAdminUsers) SetReq(req request.SysAdminUsersRequest) *SysAdminUsers {
	svc.req = req
	return svc
}

// SetReqForm 设置参数
func (svc *SysAdminUsers) SetReqForm(req request.SysAdminUsersRequestForm) *SysAdminUsers {
	req.SysAdminUsers.Username = req.Username
	req.SysAdminUsers.Password = req.Password
	req.SysAdminUsers.NickName = req.NickName

	svc.reqForm = req
	return svc
}

// SetReqLoginForm 设置参数
func (svc *SysAdminUsers) SetReqLoginForm(req request.SysAdminUsersLoginRequestForm) *SysAdminUsers {
	svc.reqLoginForm = req
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
				"username": item.Username,
				"nickname": item.NickName,
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
