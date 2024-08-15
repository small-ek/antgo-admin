package routes

import (
	"github.com/gin-gonic/gin"
	"server/app/http/api"
)

func IndexRoute(route *gin.RouterGroup) {
	IndexController := new(api.IndexController)
	route.GET("/", IndexController.Index)

	AuthController := api.NewAuthController()
	route.GET("captcha", AuthController.Captcha) //验证码
	route.POST("login", AuthController.Login)    //登录

}
