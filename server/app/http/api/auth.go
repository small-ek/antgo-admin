package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/small-ek/antgo/os/config"
	"server/app/entity/request"
	"server/app/entity/vo"
	"server/app/service"
)

// 使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type AuthController struct {
	vo.Base
	SysAdminUsersService *service.SysAdminUsers
}

func NewAuthController() *AuthController {
	return &AuthController{
		SysAdminUsersService: service.NewSysAdminUsersService(),
	}
}

//	@Tags			系统
//	@Summary		获取验证码
//	@Accept			json
//	@Produce		json
//	@Success		0	{object}	response.Write{data=vo.ResponseCaptcha}
//	@Failure		500	{object}	response.Write
//	@Router			/captcha [get] 路由
//
// Captcha
func (ctrl *AuthController) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(config.GetInt("captcha.height"), config.GetInt("captcha.width"), config.GetInt("captcha.length"), 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // 使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, _, err := cp.Generate(); err != nil {
		ctrl.Fail(c, "ERROR", err)
	} else {
		ctrl.Success(c, "SUCCESS", vo.ResponseCaptcha{
			Id:  id,
			Pic: b64s,
		})
	}
}

//	@Tags			系统
//	@Summary		用户登录
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminUsersLoginRequestForm true "登录参数"
//	@Success		0	{object}	response.Write
//	@Failure		1	{object}	response.Write
//	@Router			/login [post]
//
// Login 登录
func (ctrl *AuthController) Login(c *gin.Context) {
	var req request.SysAdminUsersLoginRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}
	if ctrl.verifyCaptcha(req.Id, req.Code) == false {
		ctrl.Fail(c, vo.CAPTCHA_ERROR, errors.New(vo.CAPTCHA_ERROR))
		return
	}
	req.DeviceId = ctrl.GetDeviceId(c)
	if result, err := ctrl.SysAdminUsersService.SetReq(req).Login(); err != nil {
		ctrl.Fail(c, vo.LOGIN_ERROR, err)
		return
	} else {
		ctrl.Success(c, vo.LOGIN_SUCCESS, result)
	}
}

// verifyCaptcha 验证验证码
func (ctrl *AuthController) verifyCaptcha(id string, VerifyValue string) bool {
	if store.Verify(id, VerifyValue, true) {
		return true
	}
	return false
}

//	@Tags			管理员用户
//	@Summary		修改用户数据
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminUsersInfoRequest true "更新参数"
//	@Success		0	{object}	response.Write
//	@Failure		1	{object}	response.Write
//	@Router			/userinfo [put]
//
// UpdateUserInfo 修改用户信息
func (ctrl *AuthController) UpdateUserInfo(c *gin.Context) {
	var req request.SysAdminUsersInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}
	req.Id = ctrl.GetUser(c).Id
	if err := ctrl.SysAdminUsersService.SetReq(req).Update(); err != nil {
		ctrl.Fail(c, vo.UPDATE_FAILED, err)
		return
	}
	ctrl.Success(c, vo.UPDATE_SUCCESS)
}

//	@Tags			管理员用户
//	@Summary		修改密码
//	@Accept			json
//	@Produce		json
//	@Param		    data body request.SysAdminUsersInfoRequest true "更新参数"
//	@Success		0	{object}	response.Write
//	@Failure		1	{object}	response.Write
//	@Router			/password [put]
//
// UpdateUserInfo 修改用户信息
func (ctrl *AuthController) UpdatePassword(c *gin.Context) {
	var req request.SysAdminUsersPasswordRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, vo.INVALID_REQUEST_PARAMETERS, err)
		return
	}
	req.Id = ctrl.GetUser(c).Id
	if err := ctrl.SysAdminUsersService.SetReq(req).UpdatePassword(); err != nil {
		ctrl.Fail(c, vo.GetUpdatePasswordErrorMsg(err), err)
		return
	}
	ctrl.Success(c, vo.UPDATE_SUCCESS)
}
