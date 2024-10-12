package routes

import (
	"github.com/gin-gonic/gin"
	"server/app/http/api"
)

func SysMenuRoute(route *gin.RouterGroup) {
	sysMenuController := api.NewSysMenuController()
	v1 := route.Group("sys-menu")
	{
		v1.GET("", sysMenuController.Index)
		v1.GET(":id", sysMenuController.Show)
		v1.DELETE(":id", sysMenuController.Delete)
		v1.POST("", sysMenuController.Create)
		v1.PUT(":id", sysMenuController.Update)
	}
}

