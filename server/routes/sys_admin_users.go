package routes

import (
	"github.com/gin-gonic/gin"
	"server/app/http/api"
	"server/app/http/middleware"
)

func SysAdminUsersRoute(route *gin.RouterGroup) {
	sysAdminUsersController := api.NewSysAdminUsersController()
	v1 := route.Group("sys-admin-users").Use(middleware.AuthJWT())
	{
		v1.GET("", sysAdminUsersController.Index)
		v1.GET(":id", sysAdminUsersController.Show)
		v1.DELETE("", sysAdminUsersController.DeleteIds)
		v1.POST("", sysAdminUsersController.Create)
		v1.PUT(":id", sysAdminUsersController.Update)
	}
}
