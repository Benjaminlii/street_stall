package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/dal"
	"street_stall/biz/domain/model"
	"street_stall/biz/util"
)

// SelectUser 查询用户信息，用于登录
func SelectUser(username string, password string) *model.User {
	user := dal.GetUserByUsernameAndPassword(username, password)
	if user == nil {
		return nil
	}
	return user
}

// SignUp 用户注册
func SignUp(username string, password string, name string, userIdentity uint, category uint) *model.User {
	db := dal.GetDB()
	// 数据库事物
	tx := db.Begin()
	defer tx.Commit()

	// user对象的构造
	user := &model.User{
		Username:     username,
		Password:     password,
		UserIdentity: userIdentity,
	}
	user = dal.InsertUser(user)

	if userIdentity == constants.USERIDENTITY_MERCHANT {
		// 商户注册
		merchant := &model.Merchant{
			UserId:   user.ID,
			Name:     name,
			Category: category,
		}
		merchant = dal.InsertMerchant(merchant)
	} else if userIdentity == constants.USERIDENTITY_VISITER {
		// 游客注册
		visitor := &model.Visitor{
			UserId: user.ID,
			Name:   name,
		}
		visitor = dal.InsertVisitor(visitor)
		if visitor == nil {
			log.Print("[service][user][SignUp] InsertVisitor fail")
			panic(constants.SYSTEM_ERROR)
		}
	}

	return user
}

// UpdateMerchantByUserId 通过userId选择商户，更新其字段信息
func UpdateMerchantByUserId(c *gin.Context, name string, category uint, introduction string) *model.Merchant {
	merchant := GetMerchantByCurrentUser(c)

	merchant.Name = name
	merchant.Category = category
	merchant.Introduction = introduction

	dal.SaveMerchant(merchant)

	return merchant
}

// GetMerchantByCurrentUser 获取当前用户对应的商户
func GetMerchantByCurrentUser(c *gin.Context) *model.Merchant {
	// 获取user
	currentUser := util.GetCurrentUser(c)

	if currentUser.UserIdentity != constants.USERIDENTITY_MERCHANT {
		log.Printf("[service][merchant][GetVisitorByCurrentUser] current user is not a merchant")
		panic(constants.AUTHORITY_ERROR)
	}

	merchant := dal.GetMerchantByUserId(currentUser.ID)

	return merchant
}

// UpdateVisitorByUserId 通过userId选择游客，更新其字段信息
func UpdateVisitorByUserId(c *gin.Context, name string, introduction string) *model.Visitor {
	visitor := GetVisitorByCurrentUser(c)

	visitor.Name = name
	visitor.Introduction = introduction

	dal.SaveVisitor(visitor)

	return visitor
}

// GetVisitorByCurrentUser 获取当前用户对应的游客
func GetVisitorByCurrentUser(c *gin.Context) *model.Visitor {
	// 获取user
	currentUser := util.GetCurrentUser(c)
	if currentUser.UserIdentity != constants.USERIDENTITY_VISITER {
		log.Printf("[service][merchant][GetVisitorByCurrentUser] current user is not a visitor")
		panic(constants.AUTHORITY_ERROR)
	}

	visitor := dal.GetVisitorByUserId(currentUser.ID)

	return visitor
}
