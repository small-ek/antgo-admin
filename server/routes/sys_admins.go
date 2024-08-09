package routes

import (
	"github.com/gin-gonic/gin"
	"server/app/http/api"
)

func SysAdminsRoute(route *gin.RouterGroup) {
	SysAdminsController := api.NewSysAdminsController()
	v1 := route.Group("/api/sys_admins")
	{
		v1.GET("", SysAdminsController.Index)
		v1.GET("/:id", SysAdminsController.Show)
		v1.DELETE("/:id", SysAdminsController.Delete)
		v1.POST("", SysAdminsController.Create)
		v1.PUT("/:id", SysAdminsController.Update)
	}
}

