package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dao"
	"street_stall/biz/domain/model"
)

// SelectUser 查询用户信息，用于登录
func SelectUser(username string, password string) *model.User {
	user := dao.GetUserByUsernameAndPassword(username, password)
	if user == nil {
		return nil
	}
	return user
}

// SignUp 用户注册
func SignUp(username string, password string, name string, userIdentity uint, category uint) *model.User {
	db := dao.GetDB()
	// 数据库事物
	tx := db.Begin()
	defer tx.Commit()

	// user对象的构造
	user := &model.User{
		Username:     username,
		Password:     password,
		UserIdentity: userIdentity,
	}
	user = dao.InsertUser(user)

	if userIdentity == constants.USERIDENTITY_MERCHANT {
		// 商户注册
		merchant := &model.Merchant{
			UserId:   user.ID,
			Name:     name,
			Category: category,
		}
		merchant = dao.InsertMerchant(merchant)
	} else if userIdentity == constants.USERIDENTITY_VISITER {
		// 游客注册
		visitor := &model.Visitor{
			UserId: user.ID,
			Name:   name,
		}
		visitor = dao.InsertVisitor(visitor)
		if visitor == nil {
			log.Print("[service][user][SignUp] InsertVisitor fail")
			panic(errors.SYSTEM_ERROR)
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

	dao.SaveMerchant(merchant)

	return merchant
}

// UpdateVisitorByUserId 通过userId选择游客，更新其字段信息
func UpdateVisitorByUserId(c *gin.Context, name string, introduction string) *model.Visitor {
	visitor := GetVisitorByCurrentUser(c)

	visitor.Name = name
	visitor.Introduction = introduction

	dao.SaveVisitor(visitor)

	return visitor
}
