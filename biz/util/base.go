package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/domain/model"
)

// GetCurrentUser 从context中获取登录状态
func GetCurrentUser(c *gin.Context) *model.User {
	currentUserInterFace, isOk := c.Get(constants.CURRENT_USER)
	if !isOk {
		log.Printf("[system][user][GetCurrentUser] current user is not exist")
		panic(constants.NO_LOGIN_ERROR)
	}
	currentUser, isOk := currentUserInterFace.(*model.User)
	if !isOk {
		panic(constants.SYSTEM_ERROR)
	}
	return currentUser
}
