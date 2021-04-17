package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"street_stall/biz/constants"
	"street_stall/biz/constants/errors"
	"street_stall/biz/service"
	"street_stall/biz/util"
)

// Reserve 预约摊位
func Reserve(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][location][Reserve] request type error, err:%s", err)
		panic(err)
	}
	placeIdStr, havePlaceId := param["place_id"]
	locationIdStr, haveLocationId := param["location_id"]
	reserveTimeStr, haveReserveTime := param["reserve_time"]
	comment, haveComment := param["comment"]
	if !(havePlaceId && haveLocationId && haveReserveTime && haveComment) {
		log.Printf("[service][location][Reserve] has nil in placeId, locationId, reserveTime and comment")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	placeId := util.StringToUInt(placeIdStr)
	locationId := util.StringToUInt(locationIdStr)
	reserveTime := util.StringToUInt(reserveTimeStr)

	order := service.Reserve(c, placeId, locationId, reserveTime, comment)
	log.Printf("[service][location][Reserve] reserve location success, placeId:%d, locationId:%d, merchantId:%d",
		placeId, locationId, order.MerchantId)

	// 设置请求响应
	respMap := map[string]interface{}{}
	c.Set(constants.DATA, respMap)
}

// GetMerchantsInfoByNameAndPlaceId 通过当前时刻商户名称获取商户信息，商户分类可选
func GetMerchantsInfoByNameAndPlaceId(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][location][GetMerchantsInfoByNameAndPlaceId] request type error, err:%s", err)
		panic(err)
	}
	placeIdStr, havePlaceId := param["place_id"]
	categoryStr, haveCategory := param["category"]
	merchantName, haveMerchantName := param["merchant_name"]
	if !(havePlaceId && haveCategory && haveMerchantName) {
		log.Printf("[service][location][GetMerchantsInfoByNameAndPlaceId] has nil in placeId, category and merchantName")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	placeId := util.StringToUInt(placeIdStr)
	category := util.StringToUInt(categoryStr)

	ans := service.GetMerchantsInfoByNameAndPlaceId(c, placeId, merchantName, category)

	// 设置请求响应
	respMap := map[string]interface{}{
		"merchants": ans,
	}
	c.Set(constants.DATA, respMap)
}
