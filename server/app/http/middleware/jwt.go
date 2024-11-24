package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/jwt"
	"github.com/small-ek/antgo/utils/response"
	"server/app/dao"
	"server/app/entity/models"
	"server/app/entity/vo"
	"strings"
)

// 常量定义
const (
	UnauthorizedCode = 401
)

func AuthJWT() gin.HandlerFunc {
	newJwt := jwt.New().SetPublicKey(conv.Bytes(config.GetString("jwt.public_key")))

	return func(c *gin.Context) {
		var token = c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_ERROR, UnauthorizedCode))
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		getAuthorization, err := newJwt.Decode(token)
		if err != nil {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_ERROR, UnauthorizedCode, err.Error()))
			return
		}
		if c.GetHeader("device-id") != getAuthorization["device-id"] {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_DEVICE_ERROR, UnauthorizedCode))
			return
		}

		getUser := getUserByUser(conv.Int(getAuthorization["id"]))
		if getUser.Id == 0 {
			c.AbortWithStatusJSON(UnauthorizedCode, response.Fail(vo.AUTH_USER_ERROR, UnauthorizedCode))
			return
		}
		c.Set("user", getUser)
		c.Set("user_id", getUser.Id)
		c.Next()
	}
}

// getUserByUser 通过 ID 获取用户信息
func getUserByUser(userID int) models.SysAdminUsers {
	user := dao.NewSysAdminUsersDao().GetById(userID)
	return user
}
