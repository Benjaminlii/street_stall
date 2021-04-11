package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/domain/model"
	"street_stall/biz/drivers"
	"time"
)

// CheckUserLoginMiddleware 用户登录状态检测中间件
func CheckUserLoginMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		// 从headers获取token
		token := c.Request.Header["Token"][0]
		if token == "" {
			log.Printf("[system][CheckUserLoginMiddleware] no token")
			c.Abort()
			c.JSON(http.StatusOK, errors.NO_TOKEN_ERROR.ChangeToResp(nil))
			return
		}

		// 去redis中根据Token获取用户信息
		userJson, err := drivers.RedisClient.Get(constants.REDIS_USER_TOKEN_PRE + token).Result()
		if err != nil {
			log.Printf("[system][CheckUserLoginMiddleware] user id wrong, token:%s", token)
			c.Abort()
			c.JSON(http.StatusOK, errors.TOKEN_WRONG_ERROR.ChangeToResp(nil))
			return
		}
		user := &model.User{}
		err = json.Unmarshal([]byte(userJson), user)
		if err != nil {
			log.Printf("[system][CheckUserLoginMiddleware] user id wrong, token:%s", token)
			c.Abort()
			c.JSON(http.StatusOK, errors.TOKEN_WRONG_ERROR.ChangeToResp(nil))
			return
		}

		// 延长登录状态失效时间
		drivers.RedisClient.Set(constants.REDIS_USER_TOKEN_PRE+token, userJson, time.Hour*24*3)

		c.Set(constants.CURRENT_USER, user)
		log.Printf("[system][CheckUserLoginMiddleware] get current user success, username:%s", user.Username)
		c.Next()
	}
}
