package api

import (
	"github.com/gin-gonic/gin"
	"server/app/entity/vo"
)

type SystemController struct {
	vo.Base
}

// Index
func (ctrl *SystemController) Index(c *gin.Context) {
	c.String(200, "Hello AntGo")
}
