package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/dao"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// GetMerchant 得到当前商户信息
func GetMerchant(c *gin.Context) {
	defer util.SetResponse(c)

	merchant := service.GetMerchantByCurrentUser(c)

	// 设置请求响应
	user := dao.GetUserById(merchant.UserId)
	if user == nil {
		panic(errors.SYSTEM_ERROR)
	}
	respMap := map[string]interface{}{
		"name":         merchant.Name,
		"introduction": merchant.Introduction,
		"user": map[string]interface{}{
			"username":      user.Username,
			"user_identity": user.UserIdentity,
		},
	}

	c.Set(constants.DATA, respMap)
}

// UpdateMerchant 商户信息维护
func UpdateMerchant(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][UpdateMerchant] request type error, err:%s", err)
		panic(err)
	}
	name, haveName := param["name"]
	categoryStr, haveCategory := param["category"]
	introduction, haveIntroduction := param["info"]
	if !(haveName && haveCategory && haveIntroduction) {
		log.Print("[service][user][UpdateMerchant] has nil in name, category and introduction")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	category := util.StringToUInt(categoryStr)

	// 更新商户信息
	merchant := service.UpdateMerchantByUserId(c, name, category, introduction)

	// 设置请求响应
	user := dao.GetUserById(merchant.UserId)
	if user == nil {
		panic(errors.SYSTEM_ERROR)
	}
	respMap := map[string]interface{}{
		"name":         merchant.Name,
		"category":     merchant.Category,
		"introduction": merchant.Introduction,
		"user": map[string]interface{}{
			"username":      user.Username,
			"user_identity": user.UserIdentity,
		},
	}

	c.Set(constants.DATA, respMap)
}

// GetMerchantByLocationId 根据摊位id获取当前位置上商户的基础信息
func GetMerchantByLocationId(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][merchant][GetMerchantByLocationId] request type error, err:%s", err)
		panic(err)
	}
	placeIdStr, havePlace := param["place_id"]
	numberOfPlaceStr, haveNumberOfPlace := param["number_of_place"]
	if !(havePlace && haveNumberOfPlace) {
		log.Print("[service][merchant][GetMerchantByLocationId] has nil in placeId and numberOfPlace")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	placeId := util.StringToUInt(placeIdStr)
	numberOfPlace := util.StringToUInt(numberOfPlaceStr)

	// 更新商户信息
	ans := service.GetMerchantByLocationId(c, placeId, numberOfPlace)

	// 设置请求响应
	c.Set(constants.DATA, ans)
}
