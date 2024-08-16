package api

import (
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
//	@Success		200	{object}	response.Write{data=vo.ResponseCaptcha}
//	@Failure		500	{object}	response.Write
//	@Router			/captcha [get] 路由
//
// Captcha
func (ctrl *AuthController) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(config.GetInt("captcha.height"), config.GetInt("captcha.width"), config.GetInt("captcha.length"), 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // 使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, _, err := cp.Generate(); err != nil {
		ctrl.Fail(c, "ERROR", err.Error())
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
//	@Success		200	{object}	response.Write
//	@Failure		422	{object}	response.Write
//	@Router			/login [post]
//
// Login 登录
func (ctrl *AuthController) Login(c *gin.Context) {
	var req request.SysAdminUsersLoginRequestForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Fail(c, "INVALID_REQUEST_PARAMETERS", err.Error())
		return
	}
	if ctrl.verifyCaptcha(req.Id, req.Code) == false {
		ctrl.Fail(c, "CAPTCHA_ERROR")
		return
	}

	if result, err := ctrl.SysAdminUsersService.SetReqLoginForm(req).Login(); err != nil {
		ctrl.Fail(c, "LOGIN_ERROR", err.Error())
		return
	} else {
		ctrl.Success(c, "SUCCESS", result)
	}
}

// verifyCaptcha 验证验证码
func (ctrl *AuthController) verifyCaptcha(id string, VerifyValue string) bool {
	if store.Verify(id, VerifyValue, true) {
		return true
	}
	return false
}
