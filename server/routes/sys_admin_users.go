package routes

import (
	"github.com/gin-gonic/gin"
	"server/app/http/api"
)

func SysAdminUsersRoute(route *gin.RouterGroup) {
	sysAdminUsersController := api.NewSysAdminUsersController()
	v1 := route.Group("sys-admin-users")
	{
		v1.GET("", sysAdminUsersController.Index)
		v1.GET(":id", sysAdminUsersController.Show)
		v1.DELETE(":id", sysAdminUsersController.Delete)
		v1.POST("", sysAdminUsersController.Create)
		v1.PUT(":id", sysAdminUsersController.Update)
	}
}
