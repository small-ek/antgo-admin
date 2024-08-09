package routes

import (
	"github.com/gin-gonic/gin"
	"server/app/http/api"
)

func IndexRoute(route *gin.RouterGroup) {
	IndexController := new(api.IndexController)
	route.GET("/", IndexController.Index)

	NewSystemController := api.NewSystemController()
	route.GET("captcha", NewSystemController.Captcha) //验证码
}
