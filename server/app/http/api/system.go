package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/small-ek/antgo/os/config"
	"server/app/entity/vo"
)

// 使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type SystemController struct {
	vo.Base
}

func NewSystemController() *SystemController {
	return &SystemController{}
}

//	@Summary		获取验证码
//	@Description	获取验证码相关数据
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Param		    data query request.ReportRequest true "分页获取列表"
//	@Success		200	{array}		response.Write
//	@Failure		500	{object}	response.Write
//	@Router			/captcha [get] 路由
//
// Captcha
func (ctrl *SystemController) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(config.GetInt("captcha.height"), config.GetInt("captcha.width"), config.GetInt("captcha.length"), 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // 使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, _, err := cp.Generate(); err != nil {
		ctrl.Fail(c, "验证码获取失败", err.Error())
	} else {
		ctrl.Success(c, "success", vo.ResponseCaptcha{
			Id:  id,
			Pic: b64s,
		})
	}
}
