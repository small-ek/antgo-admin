package routes

import (
	"server/app/http/api"
	"github.com/gin-gonic/gin"
)

func IndexRoute(route *gin.RouterGroup) {
	IndexController := new(api.IndexController)
	route.GET("/", IndexController.Index)
}

