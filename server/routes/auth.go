package routes

import (
	"github.com/gin-gonic/gin"
	"server/app/http/api"
	"server/app/http/middleware"
)

func AuthRoute(route *gin.RouterGroup) {
	AuthController := api.NewAuthController()
	route.GET("captcha", AuthController.Captcha) //验证码
	route.POST("login", AuthController.Login)    //登录
	v1 := route.Group("auth").Use(middleware.AuthJWT())
	{
		v1.PUT("userinfo", AuthController.UpdateUserInfo) //更新用户信息
	}
}
