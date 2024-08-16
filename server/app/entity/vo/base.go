package vo

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/utils/response"
	"go.uber.org/zap"
)

type Base struct {
	Code int `json:"code"` //Status Code
}

// Success 成功返回
func (b *Base) Success(c *gin.Context, msg string, data ...interface{}) {
	code := 200
	if b.Code != 0 {
		code = b.Code
	}

	c.SecureJSON(200, response.Success(msg, code, data...))
}

// Fail 错误返回
func (b *Base) Fail(c *gin.Context, msg string, err ...string) {
	if len(err) > 0 {
		alog.Write.Debug("Return error", zap.Any("error", err))
	}
	code := 422
	if b.Code != 0 {
		code = b.Code
	}

	c.SecureJSON(200, response.Fail(msg, code, err...))

}

// SetCode 修改状态
func (b *Base) SetCode(code int) *Base {
	b.Code = code
	return b
}

// Page 分页数据
func (b *Base) Page(total int64, list interface{}) response.Page {
	return response.Page{
		Total: total,
		Items: list,
	}
}
