package vo

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/utils/response"
	"go.uber.org/zap"
)

type Base struct {
	Status int `json:"status"` //Status Code
}

// Success 成功返回
func (b *Base) Success(c *gin.Context, msg string, data ...interface{}) {
	if b.Status == 0 {
		b.Status = 200
	}
	c.SecureJSON(200, response.Success(msg, b.Status, data...))
}

// Fail 错误返回
func (b *Base) Fail(c *gin.Context, msg string, err ...string) {
	if len(err) > 0 {
		alog.Write.Debug("Return error", zap.Any("error", err))
	}
	if b.Status == 0 {
		b.Status = 422
	}
	c.SecureJSON(200, response.Fail(msg, b.Status, err...))

}

// SetStatus 修改状态
func (b *Base) SetStatus(status int) *Base {
	b.Status = status
	return b
}

// Page 分页数据
func (b *Base) Page(total int64, list interface{}) response.Page {
	return response.Page{
		Total: total,
		List:  list,
	}
}
