package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/jwt"
	"github.com/small-ek/antgo/utils/response"
	"server/app/dao"
	"server/app/entity/models"
)

func AuthJWT() gin.HandlerFunc {
	var j = jwt.New().SetPublicKey(conv.Bytes(config.GetString("jwt.public_key")))

	return func(c *gin.Context) {
		var token = c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, response.Fail("用户不存在,请重新登录", 401))
			return
		}

		getAuthorization, err := j.Decode(token)
		if err != nil {
			c.AbortWithStatusJSON(401, response.Fail("error", 401, err.Error()))
			return
		}

		getUser := dao.NewSysAdminUsersDao().GetById(conv.Int(getAuthorization["id"]))
		if getUser.Id == 0 {
			c.AbortWithStatusJSON(401, response.Fail("用户不存在,请重新登录", 401))
			return
		}
		c.Set("user", getUser)
		c.Set("user_id", getUser.Id)
		c.Next()
	}
}

// GetUser 获取当前用户
func GetUser(c *gin.Context) models.SysAdminUsers {
	userModel, ok := c.MustGet("user").(models.SysAdminUsers)
	if !ok {
		return models.SysAdminUsers{}
	}
	return userModel
}
