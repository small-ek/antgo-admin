package api

import (
	"github.com/gin-gonic/gin"
	"server/app/entity/vo"
)

type IndexController struct {
	vo.Base
}

// Index
func (ctrl *IndexController) Index(c *gin.Context) {
	c.String(200, "Hello AntGo")
}
