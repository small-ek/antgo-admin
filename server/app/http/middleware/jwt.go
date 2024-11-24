package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/jwt"
	"github.com/small-ek/antgo/utils/response"
	"server/app/dao"
	"server/app/entity/vo"
	"strings"
)

func AuthJWT() gin.HandlerFunc {
	newJwt := jwt.New().SetPublicKey(conv.Bytes(config.GetString("jwt.public_key")))

	return func(c *gin.Context) {
		var token = c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, response.Fail(vo.AUTH_ERROR, 401))
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		getAuthorization, err := newJwt.Decode(token)
		if err != nil {
			c.AbortWithStatusJSON(401, response.Fail(vo.AUTH_ERROR, 401, err.Error()))
			return
		}

		getUser := dao.NewSysAdminUsersDao().GetById(conv.Int(getAuthorization["id"]))
		if getUser.Id == 0 {
			c.AbortWithStatusJSON(401, response.Fail(vo.AUTH_USER_ERROR, 401))
			return
		}
		c.Set("user", getUser)
		c.Set("user_id", getUser.Id)
		c.Next()
	}
}
