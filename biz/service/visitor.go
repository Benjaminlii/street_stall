package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dao"
	"street_stall/biz/domain/model"
	"street_stall/biz/util"
)

// GetVisitorByCurrentUser 获取当前用户对应的游客
func GetVisitorByCurrentUser(c *gin.Context) *model.Visitor {
	// 获取user
	currentUser := util.GetCurrentUser(c)
	if currentUser.UserIdentity != constants.USERIDENTITY_VISITER {
		log.Printf("[service][visitor][GetVisitorByCurrentUser] current user is not a visitor")
		panic(errors.AUTHORITY_ERROR)
	}

	visitor := dao.GetVisitorByUserId(currentUser.ID)

	return visitor
}
